package main

import (
	"fmt"

	"practice/chatx"
	"practice/chatx_demo/config"
)

func main() {
	// 调用net包中的dial 传入ip 端口 进行拨号连接，通过三次握手之后获取到conn
	chat, err := chatx.NewChatClient(config.Server_NetWorkType, config.Server_Address)
	if err != nil {
		fmt.Println("Client create conn error err:", err)
	}
	defer chat.Close()
	chat.Accept()
	fmt.Println("client")
	tm := make(chan string, 1)
	<-tm
}
