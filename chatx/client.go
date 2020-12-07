package chatx

import (
	"fmt"
	"practice/lib"
)

func (chat ChatClient) Accept() {
	//往服务端传递消息
	Write(chat.Conn, "客户连接")

	go chat.readServer()
	go chat.writeServer()
}

func (chat ChatClient) Close() {
	chat.Conn.Close()
}

func (chat ChatClient) readServer() {
	for {
		//读取服务端返回的消息
		str, err := Read(chat.Conn)

		if err == nil {
			fmt.Println(str)
		} else {
			lib.CheckError(err)
		}
	}
}

func (chat ChatClient) writeServer() {
	var msg string
	for {
		fmt.Scanln(&msg)
		Write(chat.Conn, "Scanln:"+msg)
	}
}
