package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	go func() {
		goid := GetGoid()

		fmt.Println(goid)
	}()

	time.Sleep(2 * time.Second)

}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)
	fmt.Println("the stk", stk)
	fmt.Println("------------------------------")
	fmt.Println("the field", strings.Fields(stk))
	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
