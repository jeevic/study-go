package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var rwMux sync.RWMutex

	go func() {
		fmt.Println("g1 start read lock")
		time.Sleep(2 * time.Second)
		rwMux.RLock()
		fmt.Println("g1 in read lock")
		time.Sleep(2 * time.Second)
		rwMux.RUnlock()
		fmt.Println("g1 over read lock")
	}()

	go func() {
		fmt.Println("g2 start read lock")
		time.Sleep(2 * time.Second)
		rwMux.RLock()
		fmt.Println("g2 in read lock")
		time.Sleep(2 * time.Second)
		rwMux.RUnlock()
		fmt.Println("g2 over read lock")
	}()

	go func() {
		fmt.Println("g3 start read lock")
		time.Sleep(6 * time.Second)
		rwMux.Lock()
		fmt.Println("g3 in read lock")
		time.Sleep(1 * time.Second)
		rwMux.Unlock()
		fmt.Println("g3 over read lock")
	}()

	time.Sleep(10 * time.Second)

}
