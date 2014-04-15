package push

import (
	"io"
	"io/ioutil"
	"bufio"
	"strings"
	"net"
	"net/textproto"
	"regexp"
	"encoding/base64"
)

type PushHandler interface{
	Handle(src io.Reader)
}

var filterCommand = regexp.MustCompile(`^([^ ]+)(\w+(.*))?$`)

func HandleConnection(conn net.Conn,ph PushHandler){
	protoc := textproto.NewConn(conn)
	defer protoc.Close()
	for {
		str,e := protoc.ReadLine()
		if e!=nil { return }
		cmd := filterCommand.FindStringSubmatch(str)
		if cmd==nil { return }
		switch cmd[1]{
		case "PUSH":
			r := protoc.DotReader()
			r = decodeBase64Lines(r)
			ph.Handle(r)
		}
	}
}

func decodeBase64Lines(src io.Reader) io.Reader{
	r,w := io.Pipe()
	bsrc := bufio.NewReader(src)
	go func(){
		defer w.Close()
		for{
			data,err := bsrc.ReadBytes('\n')
			if len(data)>0 {
				data,e := base64.StdEncoding.DecodeString(strings.Trim(string(data),"\r\n\t "))
				if e!=nil {
					io.Copy(ioutil.Discard,src)
					return
				}
				w.Write(data)
			}
			if err!=nil { break }
		}
	}()
	return r
}

