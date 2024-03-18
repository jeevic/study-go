package main

import (
	"fmt"
	"sync"
	"time"
)

var s []int

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i <= 10000; i++ {

			time.Sleep(100 * time.Microsecond)
			s = append(s, i)
		}
		fmt.Println("over add int")
		wg.Done()
	}()

	go func() {
		time.Sleep(1)
		for k, v := range s {
			fmt.Printf("k: %d v: %d \n", k, v)
		}
		fmt.Println("over output ")

		wg.Done()
	}()

	wg.Wait()

}
