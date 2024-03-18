package main

import (
	"log"
	"runtime/debug"
)

func Test(f func() int) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			println("----------------------------------")
			log.Println(err)
		}
	}()
	x := f()
	println(x)
}
func main() {
	a := 0
	Test(func() int {
		return 100 / a
	})
}
