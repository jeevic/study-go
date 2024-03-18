package main

import "fmt"

func main() {
	var a int
	if true {
		a := 1
		a += 1
	}

	fmt.Println(a)
}
