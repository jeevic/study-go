package ants

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

var n = 100000
var curMem uint64

func TestAntsPoolWithFunc(t *testing.T) {
	var wg sync.WaitGroup
	p, _ := NewPoolWithFunc(AntsSize, func(i interface{}) error {
		demoPoolFunc(i)
		wg.Done()
		return nil
	})
	defer p.Release()

	for i := 0; i < n; i++ {
		wg.Add(1)
		p.Serve(Param)
	}
	wg.Wait()
	t.Logf("pool with func, running workers number:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d", curMem)
}

func TestNoPool(t *testing.T) {
	t.Logf("log start")

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}

	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)

}

func TestAntsPool(t *testing.T) {
	defer Release()
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		Submit(func() error {
			demoFunc()
			wg.Done()
			return nil
		})
	}
	wg.Wait()

	t.Logf("pool, capacity:%d", Cap())
	t.Logf("pool, running workers number:%d", Running())
	t.Logf("pool, free workers number:%d", Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestCodeCov(t *testing.T) {
	_, err := NewTimingPool(-1, -1)
	t.Log(err)
	_, err = NewTimingPool(1, -1)
	t.Log(err)
	_, err = NewTimingPoolWithFunc(-1, -1, demoPoolFunc)
	t.Log(err)
	_, err = NewTimingPoolWithFunc(1, -1, demoPoolFunc)
	t.Log(err)

	p0, _ := NewPool(AntsSize)
	defer p0.Submit(demoFunc)
	defer p0.Release()
	t.Logf("pool, capacity:%d", p0.Capacity())
	t.Logf("pool, running workers number:%d", p0.Running())
	t.Logf("pool, free workers number:%d", p0.Free())

	for i := 0; i < n; i++ {
		p0.Submit(demoFunc)
	}
	t.Logf("pool, capacity:%d", p0.Capacity())

	t.Logf("pool, running workers number:%d", p0.Running())
	t.Logf("pool, free workers number:%d", p0.Free())
	p0.ReSize(AntsSize)
	p0.ReSize(AntsSize / 2)
	t.Logf("pool, after resize, capacity:%d, running:%d", p0.Capacity(), p0.Running())
	time.Sleep(1 * time.Second)
	t.Logf("pool, after resize, capacity:%d, running:%d", p0.Capacity(), p0.Running())

	p, _ := NewPoolWithFunc(TestSize, demoPoolFunc)
	defer p.Serve(Param)
	defer p.Release()
	for i := 0; i < n; i++ {
		p.Serve(Param)
	}
	time.Sleep(DefaultCleanIntervalTime * time.Second)
	t.Logf("pool with func, capacity:%d", p.Capacity())
	t.Logf("pool with func, running workers number:%d", p.Running())
	t.Logf("pool with func, free workers number:%d", p.Free())
	p.ReSize(TestSize)
	p.ReSize(AntsSize)
	t.Logf("pool with func, after resize, capacity:%d, running:%d", p.Capacity(), p.Running())
}
