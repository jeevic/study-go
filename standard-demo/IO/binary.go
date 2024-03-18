package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func checkError(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func write() {
	f, err := os.Create("test.dat")
	checkError(err)

	defer func() {
		f.Sync()
		f.Close()
	}()

	var i int32 = 0x1234
	checkError(binary.Write(f, binary.LittleEndian, i))

	var d float32 = 0.1234
	checkError(binary.Write(f, binary.LittleEndian, d))

	var s string = "Hello, 中国!"
	checkError(binary.Write(f, binary.LittleEndian, int32(len(s))))
	_, err = f.WriteString(s)

}

func read() {
	f, err := os.Open("test.dat")
	checkError(err)

	defer f.Close()

	var i int32
	checkError(binary.Read(f, binary.LittleEndian, &i))

	var d float32
	checkError(binary.Read(f, binary.LittleEndian, &d))

	var l int32
	s := make([]byte, l)
	_, err = f.Read(s)

	fmt.Printf("%#x; %f; %s\n", i, d, string(s))

}

func main() {
	write()
	read()
}
