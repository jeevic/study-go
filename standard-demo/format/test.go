package main

import (
	"fmt"
)

type student struct {
	age  uint8
	name string
}

func main() {
	/*a := student{age : 1, name : "caojianwei"}

	var wg sync.WaitGroup
	wg.Done()
	wg.Wait()


	fmt.Println(reflect.TypeOf(a).String())
	*/

	str := "Go爱好者 "
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))

}
