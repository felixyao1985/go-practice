package chatx

import (
	"fmt"
	"net"
	"time"
)

func (chat ChatServer) Accept() {
	go chat.scanln()    //监听键盘输入
	go chat.Broadcast() //开启广播
	for {
		conn, err := chat.TCPListener.Accept() //等待端口连接。

		if err != nil {
			continue
		}
		key := parsemd5(conn.RemoteAddr().String())
		chat.Coons[key] = ClintCoon{
			Conn:     conn,
			NickName: key,
			Close:    make(chan string),
			Id:       key,
		}
		go chat.handleClient(key)
	}
	defer chat.TCPListener.Close()
}

func (chat ChatServer) Close() {
	chat.TCPListener.Close()
}

func (chat ChatServer) handleClient(key string) {
	if c, ok := chat.Coons[key]; ok {
		// 设置读取超时时间
		chat.Coons[key].Conn.SetReadDeadline(time.Now().Add(20 * time.Minute))
		defer chat.Coons[key].Conn.Close()

		Write(chat.Coons[key].Conn, "服务器分配名称 :"+key)
		go chat.readCline(chat.Coons[key])
		//go chat.writeCline(chat.Coons[key])

		<-c.Close
	}
}

func (chat ChatServer) readCline(cc ClintCoon) {
	for {
		// 调用公用方法read 获取客户端传过来的消息。
		str, err := Read(cc.Conn)
		if err != nil {
			fmt.Println("config.Read :", err)
			cc.Close <- "end"
			break
		}
		fmt.Println(cc.NickName, cc.Conn.RemoteAddr(), str)
		chat.send(str, cc.Id)
	}
}

func (chat ChatServer) writeCline(cc ClintCoon) {
	for {
		select {
		case msg1 := <-chat.Msg:
			_, err := Write(cc.Conn, "Scanln:"+msg1)
			if net.ErrWriteToConnected == err {
				break
			}
		}
	}
}

func (chat ChatServer) send(msg string, key string) {
	var nickName string
	if key == "" {
		nickName = "服务器"
	} else {
		if c, ok := chat.Coons[key]; ok {
			nickName = c.NickName
		} else {
			return
		}
	}

	for _, item := range chat.Coons {
		if item.Id != key {
			_, err := Write(item.Conn, nickName+":"+msg)
			if net.ErrWriteToConnected == err {
				break
			}
		}

	}
}

func (chat ChatServer) Broadcast() {
	for {
		select {
		case msg := <-chat.Msg:
			chat.send(msg, "")
		}
	}
}

func (chat ChatServer) scanln() {
	for {
		m := ""
		fmt.Scanln(&m)
		fmt.Println("服务器输入了" + m)
		chat.Msg <- m
	}
}
