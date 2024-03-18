package ants

import (
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type pf func(interface{}) error

type PoolWithFunc struct {
	//容量
	capacity int32

	//正在执行数量
	running int32

	expiryDuration time.Duration

	//释放信号
	freeSignal chan sig

	//存储可用worker
	workers []*WorkerWithFunc

	workerPool sync.Pool

	//池子关闭标志
	release chan sig

	lock sync.Mutex

	poolFunc pf

	once sync.Once
}

func NewPoolWithFunc(size int, f pf) (*PoolWithFunc, error) {
	return NewTimingPoolWithFunc(size, DefaultCleanIntervalTime, f)
}

// 定时生成
func NewTimingPoolWithFunc(size, expiry int, f pf) (*PoolWithFunc, error) {
	if size < 0 {
		return nil, ErrInvalidPoolSize
	}

	if expiry < 0 {
		return nil, ErrInvalidPoolExpiry
	}

	p := &PoolWithFunc{
		capacity:       int32(size),
		freeSignal:     make(chan sig, math.MaxInt32),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
		poolFunc:       f,
	}

	//定期清理净化
	go p.periodicallyPurge()

	return p, nil
}

func (p *PoolWithFunc) Serve(args interface{}) error {
	if len(p.release) > 0 {
		return ErrorPoolClosed
	}
	p.getWorker().args <- args
	return nil
}

// 获取一个worker
func (p *PoolWithFunc) getWorker() *WorkerWithFunc {
	var w *WorkerWithFunc

	waiting := false

	//判断正在运行的goroutines 是否有空闲的
	p.lock.Lock()
	workers := p.workers
	n := len(workers) - 1
	if n < 0 {
		//判断是否超过容量
		waiting = p.Running() >= p.Capacity()
	} else {
		w = workers[n]
		workers[n] = nil
		p.workers = workers[:n]
	}
	p.lock.Unlock()

	if waiting {
		<-p.freeSignal

		for {
			p.lock.Lock()
			workers = p.workers
			l := len(workers) - 1
			if l < 0 {
				p.lock.Unlock()
				continue
			}

			w = workers[l]
			workers[l] = nil
			p.workers = workers[:l]
			p.lock.Unlock()

			break
		}

	} else if w == nil {
		w = &WorkerWithFunc{
			pool: p,
			args: make(chan interface{}, 1),
		}
		w.run()
		p.incrRunning()
	}

	return w

}

// 将worker 放入协程池中 循环利用
func (p *PoolWithFunc) putWorker(worker *WorkerWithFunc) {
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
	p.freeSignal <- sig{}

}

// 定期清理
func (p *PoolWithFunc) periodicallyPurge() {
	heartbeat := time.NewTicker(p.expiryDuration)

	for range heartbeat.C {
		currentTime := time.Now()

		p.lock.Lock()

		idleWorkers := p.workers

		if len(idleWorkers) == 0 && p.Running() == 0 && len(p.release) > 0 {
			p.lock.Unlock()
			return
		}

		n := -1
		for i, w := range idleWorkers {
			if currentTime.Sub(w.recycleTime) <= p.expiryDuration {
				break
			}

			n = i
			w.args <- nil
			p.workers[i] = nil

		}

		if n > -1 {
			if n >= len(idleWorkers)-1 {
				p.workers = idleWorkers[:0]
			} else {
				p.workers = idleWorkers[n+1:]
			}
		}
		p.lock.Unlock()
	}

}

// 获取正在运行worker
func (p *PoolWithFunc) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

// 增加 runing 数量
func (p *PoolWithFunc) incrRunning() {
	atomic.AddInt32(&p.running, 1)
}

// 减小runing 数量
func (p *PoolWithFunc) decRunning() {
	atomic.AddInt32(&p.running, -1)
}

// 获取容量
func (p *PoolWithFunc) Capacity() int {
	return int(atomic.LoadInt32(&p.capacity))
}

// 返回可用的 gorouting
func (p *PoolWithFunc) Free() int {
	return int(atomic.LoadInt32(&p.capacity) - atomic.LoadInt32(&p.running))
}

// 调整容量
func (p *PoolWithFunc) ReSize(size int) error {
	if size < 0 {
		return ErrInvalidPoolSize
	}
	if size == p.Capacity() {
		return nil
	}

	atomic.StoreInt32(&p.capacity, int32(size))

	diff := p.Running() - size

	//如果 running 的数量大于 size 清理掉
	if diff > 0 {
		for i := 0; i < diff; i++ {
			worker := p.getWorker()
			worker.args <- nil

		}
	}

	return nil

}

// 释放池资源
func (p *PoolWithFunc) Release() error {
	p.once.Do(func() {
		p.release <- sig{}
		p.lock.Lock()
		idleWorkers := p.workers

		for i, w := range idleWorkers {
			w.args <- nil
			p.workers[i] = nil
		}
		p.workers = nil

		p.lock.Unlock()
	})

	return nil
}
