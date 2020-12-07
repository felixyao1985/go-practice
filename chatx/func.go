package chatx

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net"
)

const (
	Delimiter = '\t'
)

func parsemd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	l := fmt.Sprintf("%x", h.Sum(nil))

	return l
}

// 往conn中写数据，可以用于客户端传输给服务端， 也可以服务端返回客户端
func Write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(Delimiter)

	return conn.Write(buffer.Bytes())
}

// 从conn中读取字节流，以上面的结束符为标记
func Read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		if _, err := conn.Read(readBytes); err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == Delimiter {
			break
		}
		buffer.WriteByte(readByte)
	}

	return buffer.String(), nil
}
