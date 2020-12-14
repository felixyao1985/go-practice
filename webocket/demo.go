package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

type WSServer struct {
	ListenAddr string
}

func (this *WSServer) handler(conn *websocket.Conn) {
	fmt.Printf("a new ws conn: %s->%s\n", conn.RemoteAddr().String(), conn.LocalAddr().String())
	var err error
	//content := make([]byte, 1024)
	//_, err = conn.Read(content)
	//fmt.Println("content ?:",string(content))
	//if err != nil {
	//	fmt.Println("content Read err:", err.Error())
	//	return
	//}
	fmt.Println(conn.Request().Header["Accept-Encoding"])
	for {
		var reply string
		err = websocket.Message.Receive(conn, &reply)
		if err != nil {
			fmt.Println("receive err:", err.Error())
			break
		}

		fmt.Println("Received from client: " + reply)
		if err = websocket.Message.Send(conn, reply); err != nil {
			fmt.Println("send err:", err.Error())
			break
		}
	}
}

func (this *WSServer) start() error {
	http.Handle("/ws", websocket.Handler(this.handler))
	fmt.Println("begin to listen")
	err := http.ListenAndServe(this.ListenAddr, nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
		return err
	}
	fmt.Println("start end")
	return nil
}

func main() {
	addr := flag.String("addr", "127.0.1.1:8888", "websocket server listen address")
	flag.Parse()
	wsServer := &WSServer{
		ListenAddr: *addr,
	}
	wsServer.start()
	fmt.Println("------end-------")
}
