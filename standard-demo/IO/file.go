package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	f, err := os.Open("D:\\webapp\\phpstudy\\WWW\\Go\\standard\\IO\\text.txt")
	if err != nil {
		print(err)
		println("error ")
		println(f)
		return
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 4)

	for {
		line, isPrefix, err := r.ReadLine()

		println(isPrefix)
		println(line)

		if isPrefix {
			println(string(line))
		} else if len(line) > 0 {
			println(string(line))
		}

		if err == io.EOF {
			break
		}
	}

}
