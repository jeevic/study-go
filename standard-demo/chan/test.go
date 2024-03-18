package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func main() {
	c := make(chan *User, 10)

	go func() {
		user := User{Name: "jeevi"}
		c <- &user
	}()
	for i := 0; i < 10; i++ {
		go func(it int) {
			for {
				select {
				case u := <-c:
					fmt.Printf("i:%d %p %#v \n", it, u, u)
				default:
					time.Sleep(time.Second * 1)
				}
			}
		}(i)
	}

	t := time.Tick(time.Second * 1)
	for {
		select {
		case <-t:
			user := User{Name: fmt.Sprintf("jeevi %s", time.Now().Format("2006-01-02 15:04:05.0000000)"))}
			c <- &user
		}
	}

}
