package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	//初始化日志
	log.SetPrefix("[ http proxy ]")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}
func main() {

	var ip = flag.String("ip", "", "listen ip")
	var port = flag.Int("port", 6666, "the port is int")

	if *ip != "" {
		ipv4 := net.ParseIP(*ip)

		if ipv4 == nil {
			fmt.Println("ip error!")
			return
		}
	}

	if *port < 1 || *port >= 65535 {
		fmt.Println("port not valid ")
		return
	}

	log.Println("http proxy start")
	listen := *ip + ":" + strconv.Itoa(*port)

	log.Println("listen address", listen)
	ln, err := net.Listen("tcp", listen)

	if err != nil {
		log.Fatalln("establish connection error ", err)

		return
	}

	defer ln.Close()

	//loop listen

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatalln("accept error", err)

			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	defer conn.Close()

	timeStart := time.Now()

	var b [1024]byte
	n, err := conn.Read(b[:])
	if err != nil {
		if err != io.EOF {
			log.Println(err)
			return
		}
	}

	/*

		CONNECT LIKE THIS
		     CONNECT www.google.com:443 HTTP/1.1
		 	 Host: www.google.com:443
			 Proxy-Connection: keep-alive
		     User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36
		HTTP
			GET http://www.flysnow.org/ HTTP/1.1
			Host: www.flysnow.org
			Proxy-Connection: keep-alive
			Upgrade-Insecure-Requests: 1
			User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36
		HTTPS
			GET  www.flysnow.org/ HTTP/1.1
			Host: www.flysnow.org
			Proxy-Connection: keep-alive
			Upgrade-Insecure-Requests: 1
			User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36
	*/

	ranNum := randNum()
	prefix := fmt.Sprintf(" [http proxy] [uid: %d]", ranNum)
	log.SetPrefix(prefix)

	var method, host, address string

	//log.Printf("binary content: %#v \n", b[:])

	log.Printf("content: %#v \n", string(b[:]))

	if bytes.IndexByte(b[:], '\n') == -1 {
		fmt.Sscanf(string(b[:]), "%s %s", &method, &host)
	} else {
		fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s %s", &method, &host)
	}

	log.Println("get method and host", method, host)

	u, err := url.Parse(host)
	if err != nil {
		// 08imgmini.eastday.com:443  error
		if !strings.HasPrefix(host, "/") {
			host = "//" + host
		}
		u, err = url.Parse(host)

		if err != nil {
			log.Fatalln("parse url error", err)
			return
		}
	}

	log.Printf("url format: %#v", u)

	if u.Opaque == "443" {
		address = u.Scheme + ":443"
	} else {
		if strings.Index(u.Host, ":") == -1 {
			address = u.Host + ":80"

		} else {
			address = u.Host
		}
	}

	log.Println("address:", address)
	//request client
	server, err := net.Dial("tcp", address)

	if err != nil {
		log.Println("link error", err)
		return
	}

	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}

	var scBuffer [512]byte
	go io.CopyBuffer(server, conn, scBuffer[:])

	var csBuffer [512]byte
	io.CopyBuffer(conn, server, csBuffer[:])

	timeEnd := time.Now()

	diff := timeEnd.Sub(timeStart).Seconds()

	log.Println("host:", host, "timestamp:", diff)

}

func randNum() int64 {

	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	randNum := rand.Int63n(9999999999)

	return randNum

}

func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "CONNECT" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "405 must CONNECT\n")
		return
	}

}
