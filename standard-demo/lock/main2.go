package main

import (
	"fmt"
	"sync"
)

func main() {

	answerReady := false
	answer := 0

	for i := 0; i < 1000000000; i++ {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			answerReady = true
			answer = 42
			wg.Done()
		}()
		go func() {
			if answerReady {
				fmt.Printf("the answer is %d \n", answer)
			}
			wg.Done()
		}()
		wg.Wait()
		answerReady = false
		answer = 0

	}

}
