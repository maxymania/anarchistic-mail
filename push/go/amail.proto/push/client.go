package push

import (
	"io"
	"bufio"
	"net"
	"net/textproto"
	"encoding/base64"
)

type PushClient struct{
	Conn *textproto.Conn
}

func NewPushClient(network, addr string) (*PushClient,error){
	c,e := textproto.Dial(network,addr)
	return &PushClient{c},e
}
func NewPushClientConn(conn net.Conn) *PushClient{
	c := textproto.NewConn(conn)
	return &PushClient{c}
}
func (conn *PushClient) Close() error{
	return conn.Conn.Close()
}
func (conn *PushClient) Handle(src io.Reader) {
	conn.Conn.PrintfLine("PUSH")
	buf := make([]byte,60)
	dest := conn.Conn.DotWriter()
	bdest := bufio.NewWriter(dest)
	for {
		n,e := src.Read(buf)
		if n>0 {
			line := base64.StdEncoding.EncodeToString(buf[:n])+"\r\n"
			bdest.WriteString(line)
		}
		if e!=nil { break }
	}
	bdest.Flush()
	dest.Close()
}

