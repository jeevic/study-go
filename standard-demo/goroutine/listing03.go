package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有goroutine 都要增加其值的变量
	counter int64

	// wg 用来等待程序结束
	wg sync.WaitGroup
)

// main 是所有Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
