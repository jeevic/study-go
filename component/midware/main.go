package main

import (
	"fmt"
)

type Handler func(int)

type Middleware func(Handler) Handler

func Chain(outer Middleware, other ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(other) - 1; i >= 0; i-- {
			next = other[i](next)
		}
		return outer(next)
	}
}

func One(i int) {
	fmt.Println("handler retrun i:", i)
}

func Log(handler Handler) Handler {
	return func(i int) {
		fmt.Println("---logging start---")
		handler(i)
		fmt.Println("---logging end---")
	}
}

func Url(handler Handler) Handler {
	return func(i int) {
		fmt.Println("---url start ---")
		handler(i)
		fmt.Println("---url end ---")
	}
}

func Path(handler Handler) Handler {
	return func(i int) {
		fmt.Println("---path start---")
		handler(i)
		fmt.Println("---path end ---")
	}
}

func main() {
	end := Chain(Log, Url, Path)
	end(One)(10)
}
