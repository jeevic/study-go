package main

import (
	"fmt"
	"net"
	"os"
)

func UdpServer(ip, port string, stop chan struct{}) {
	u, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Printf("udp server error:%s", err.Error())
		os.Exit(1)
	}
	udpConn, err := net.ListenUDP("udp", u)
	if err != nil {
		fmt.Printf("udp server error:%s", err.Error())
		os.Exit(1)
	}

	for {
		select {
		case <-stop:
			_ = udpConn.Close()
		default:
			data := make([]byte, 1024)
			_, _, err := udpConn.ReadFromUDP(data[:])
			if err != nil {
				fmt.Printf("udp server recive error:%s", err.Error())
				os.Exit(1)
			}
			/*go func(n int, addr *net.UDPAddr, data []byte) {
				// 返回数据
				fmt.Printf("Addr:%s,data:%v count:%d \n", addr, string(data[:n]), n)
				_,err := udpConn.WriteToUDP(data[:n],addr)
				if err != nil {
					fmt.Println("write to udp server failed,err:", err)
				}
			}(n,addr, data)*/
		}

	}

}

func main() {
	ip := "127.0.0.1"
	port := "15688"
	ch := make(chan struct{})
	UdpServer(ip, port, ch)
}
