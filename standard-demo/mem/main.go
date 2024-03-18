package main

import (
	"fmt"
	"syscall"
)

func main() {
	/*go setup()
	for !done {}
	print("the done flag", done)
	print(a)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
	os.Exit(0)*/

	err := syscall.Chmod(":invalid path:", 0666)
	if err != nil {
		fmt.Println(err.(syscall.Errno))
	}

}
