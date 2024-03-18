package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是所有goroutine 都要增加其值的变量
	counter int

	// wg 用来等待程序结束
	wg sync.WaitGroup

	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
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
	fmt.Printf("Final Counter: %d\\n", counter)
}

// incCounter 使用互斥锁来同步并保证安全访问，
// 增加包里counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine 进入
		// 这个临界区
		mutex.Lock()
		{
			// 捕获 counter 的值
			value := counter
			fmt.Printf("before sched: %d\\n", id)
			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()
			fmt.Printf("after sched: %d\\n", id)
			// 增加本地 value 变量的值
			value++

			// 将该值保存回counter
			counter = value
		}
		mutex.Unlock()
		// 释放锁，允许其他正在等待的goroutine
		// 进入临界区
	}
}
