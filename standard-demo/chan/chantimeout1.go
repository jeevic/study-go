package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 20)

		intChan <- 1

	}()

	ts0 := time.Now().Unix()
	fmt.Printf("time0: %v \n", time.Now())
	select {
	case e := <-intChan:
		fmt.Printf("Recived: %v\n", e)
	case <-time.NewTimer(time.Second * 5).C:
		fmt.Printf("time0: %v \n", time.Now())
		ts1 := time.Now().Unix()
		fmt.Printf("diff second: %d \n", (ts1 - ts0))
		fmt.Print("Timeout! \n")
	}

	fmt.Println("over! \n")
}
