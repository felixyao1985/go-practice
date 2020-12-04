package main

import (
	"fmt"
	"net"
	"practice/lib"
	"practice/scoket/config"
	"time"
)

var msg string

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", config.Server_Address)
	lib.CheckError(err)
	listener, err := net.ListenTCP(config.Server_NetWorkType, tcpAddr)
	lib.CheckError(err)

	go scanln()
	for {
		conn, err := listener.Accept() //等待端口连接。

		if err != nil {
			continue
		}

		go handleClient(conn)
	}
	defer listener.Close()
}

func handleClient(conn net.Conn) {
	// 设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(20 * time.Minute))
	defer conn.Close()
	go readCline(conn)
	go writeCline(conn)
	tm := make(chan string, 1)
	<-tm
}

func readCline(conn net.Conn) {
	for {
		// 调用公用方法read 获取客户端传过来的消息。
		str, err := config.Read(conn)
		if err != nil {
			fmt.Println("config.Read :", err)
			break
		}
		fmt.Println("client:", conn.RemoteAddr(), str)
	}
}

func writeCline(conn net.Conn) {
	for {
		if msg != "" {
			_, err := config.Write(conn, "Scanln:"+msg)
			msg = ""
			if net.ErrWriteToConnected == err {
				break
			}
		}
	}
}

func scanln() {
	for {
		fmt.Scanln(&msg)
		fmt.Println("服务器输入了" + msg)
	}
}
