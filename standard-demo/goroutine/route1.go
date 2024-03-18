package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(10e9)
}

func sendData(ch chan string) {
	fmt.Printf("%s ", "Washington")
	ch <- "Washington"
	fmt.Printf("%s ", "Washington")
	time.Sleep(1e9)
	ch <- "Tripoli"
	fmt.Printf("%s ", "Tripoli")
	time.Sleep(1e9)
	ch <- "London"
	fmt.Printf("%s ", "London")
	time.Sleep(1e9)
	ch <- "Beijing"
	fmt.Printf("%s ", "Beijing")
	ch <- "Tokio"
	fmt.Printf("%s ", "Tokio")
}

func getData(ch chan string) {
	var input string
	time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
