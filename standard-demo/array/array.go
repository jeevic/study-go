package main

import (
	"fmt"
	"syscall"
)

func main() {
	syscall.SIGALRM
	slice := make([]int, 3, 5)
	fmt.Println(slice[2])

}
