package main

import (
	"log"
	"net"
	"net/rpc"

	lr "github.com/jeevic/study-go/standard-demo/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(lr.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
