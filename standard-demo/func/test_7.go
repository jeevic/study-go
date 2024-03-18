package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {

			fmt.Println("err value:", err)

			err = ErrDidNotWork
		}
	}
	fmt.Println("err value end:", err)
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
