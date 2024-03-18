package main

import (
	"fmt"
	"time"
)

func main() {
	/*var a float64
	a = math.NaN()
	var b float64
	if math.IsNaN(a) {
		fmt.Println("math is Nan")
	} else {
		fmt.Println("no")
	}

	if math.IsNaN(b) {
		fmt.Println("b math is Nan")
	} else {
		fmt.Println("no")
	}*/

	/*var a int = 10

	var b int = 2

	var c int = 1

	a -= b +c
	fmt.Println(a)



	var x float64 = 5.7
	var y int = int(1.9)
	fmt.Println(y)*/

	/*go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")*/

	oldTime, _ := time.Parse("2006-01-02T15:04:05", "2017-04-01T00:00:00")

	fmt.Println(oldTime.Unix())

	fmt.Println(oldTime)

	t := []byte{'a', 'b', 'c'}
	fmt.Printf("byte1 address:%p", t)
	test(t)
}

func test(t []byte) {
	fmt.Printf("byte address:%p", t)

}
