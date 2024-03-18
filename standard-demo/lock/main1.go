package main

import (
	"fmt"
	"sync"
)

func main() {

	var a int
	a = 0
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i <= 10000; i++ {
			a++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 10000; i++ {
			a++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(a)

}
