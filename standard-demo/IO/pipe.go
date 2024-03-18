package main

import (
	"fmt"
	"io"
)

func main() {
	e := make(chan bool)

	r, w := io.Pipe()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data%d\n", i)
		}

		w.Close()
	}()

	go func() {
		for {
			var s string
			_, err := fmt.Fscanln(r, &s)
			if err == io.EOF {
				break
			}
			println(s)
		}

		e <- true
	}()

	<-e

}
