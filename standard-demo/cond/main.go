package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 3)
	max := 5
	go func(max int) { // 用于发信。
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			if mailbox == 1 {
				log.Printf("sender [%d]: check mailbox is full.", i)
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max * 2)

	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			lock.RLock()
			if mailbox == 0 {
				log.Printf("receiver1 [%d]: check mailbox is empty.", j)
				recvCond.Wait()
			}
			log.Printf("receiver1 [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver1 [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	go func(max int) { // 用于收信。
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			lock.RLock()
			if mailbox == 0 {
				log.Printf("receiver2 [%d]: check mailbox is empty.", j)
				recvCond.Wait()
			}
			log.Printf("receiver2 [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver2 [%d]: the letter has been received.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
	<-sign

}
