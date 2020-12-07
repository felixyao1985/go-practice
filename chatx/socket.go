package chatx

import (
	"net"
)

type ClintCoon struct {
	Conn     net.Conn
	NickName string
	Close    chan string
	Id       string
}
type ChatServer struct {
	TCPListener *net.TCPListener
	Coons       map[string]ClintCoon
	Msg         chan string
}

type ChatClient struct {
	Conn net.Conn
}

func NewChatServer(netType string, addr string) (ChatServer, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return ChatServer{}, err
	}
	listener, err := net.ListenTCP(netType, tcpAddr)
	if err != nil {
		return ChatServer{}, err
	}

	return ChatServer{
		TCPListener: listener,
		Msg:         make(chan string),
		Coons:       make(map[string]ClintCoon, 0),
	}, nil
}

func NewChatClient(netType string, addr string) (ChatClient, error) {
	// 调用net包中的dial 传入ip 端口 进行拨号连接，通过三次握手之后获取到conn
	conn, err := net.Dial(netType, addr)
	if err != nil {
		return ChatClient{}, err
	}
	return ChatClient{
		Conn: conn,
	}, nil
}
