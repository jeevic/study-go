package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	go func() {
		lock.RLock()
		for mailbox == 0 {
			fmt.Println(" recvie condtion wait")
			recvCond.Wait()
		}
		mailbox = 0
		lock.RUnlock()
		sendCond.Signal()

	}()

	go func() {

		lock.Lock()
		for mailbox == 1 {
			fmt.Println(" send condtion wait")
			sendCond.Wait()
		}
		mailbox = 1
		lock.Unlock()
		recvCond.Signal()

	}()

	time.Sleep(100 * time.Second)
}
