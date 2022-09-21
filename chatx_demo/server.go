package main

import (
	"practice/chatx"
	"practice/chatx_demo/config"
	"practice/lib"
)

func main() {
	chat, err := chatx.NewChatServer(config.Server_NetWorkType, config.Server_Address)

	lib.CheckError(err)

	chat.Accept()
	defer chat.Close()

	{

	}
}
