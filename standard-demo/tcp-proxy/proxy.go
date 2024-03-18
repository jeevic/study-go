package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	dstList := make(map[string]uint8)
	dstList["127.0.0.1:80"] = 1 //转发到ip列表 标记为1 代表要会结果
	go server("0.0.0.0", 8088, dstList)

	select {}
}

func server(ip string, port int32, dsts map[string]uint8) {
	//拼装地址 ip port
	addr := fmt.Sprintf("%s:%d", ip, port)

	//监听端口
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Printf("Listen地址: %s失败! 错误: %v\n", addr, err)
		return
	}
	//链接关闭
	defer listen.Close()

	fmt.Printf("Listen地址: %s 成功!\n", addr)

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		// 获取对应要转发到ip 进行转发
		for k, v := range dsts {
			go handle(conn, k, v)
		}
	}
}

// handle处理程序
func handle(sconn net.Conn, addr string, sign uint8) {
	defer sconn.Close()
	//设置tcp连接时长
	sconn.SetReadDeadline(time.Now().Add(time.Second * 200))

	dconn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}
	defer dconn.Close()

	exitChan := make(chan bool, 2)
	//从源地址往目标地址写
	go func(dc net.Conn, sc net.Conn) {
		io.Copy(dc, sc)

		exitChan <- true
	}(dconn, sconn)

	//从目标地址往源地址写
	go func(dc net.Conn, sc net.Conn, sign uint8) {
		if sign == uint8(1) {
			io.Copy(sc, dc)
		} else {
			for {
				bytes := make([]byte, 4096)
				n, _ := sc.Read(bytes) //读数据
				if n == 0 {
					break
				}
			}
		}

		exitChan <- true

	}(dconn, sconn, sign)

	<-exitChan
	<-exitChan
}
