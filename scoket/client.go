package main

import (
	"fmt"
	"net"
	"practice/lib"
	"practice/scoket/config"
)

func main() {
	// 调用net包中的dial 传入ip 端口 进行拨号连接，通过三次握手之后获取到conn
	conn, err := net.Dial(config.Server_NetWorkType, config.Server_Address)
	if err != nil {
		fmt.Println("Client create conn error err:", err)
	}
	defer conn.Close()
	//往服务端传递消息
	config.Write(conn, "客户连接")

	go readServer(conn)
	go writeServer(conn)
	tm := make(chan string, 1)
	<-tm
}

func readServer(conn net.Conn) {
	for {
		//读取服务端返回的消息
		str, err := config.Read(conn)

		if err == nil {
			fmt.Println(str)
		} else {
			lib.CheckError(err)
		}
	}
}

func writeServer(conn net.Conn) {
	var msg string
	for {
		fmt.Scanln(&msg)
		config.Write(conn, "Scanln:"+msg)
	}
}
