package vsep

import "fmt"
import "io"

func verifyLength(k []byte,req int) error {
	if k==nil { return nil }
	if len(k)==req { return nil }
	return KeyLengthError(len(k))
}

type KeyAuthenticationFailed struct{}
func (k KeyAuthenticationFailed) Error() string {
	return "Key Authentication Failed"
}

type KeyLengthError int
func (k KeyLengthError) Error() string {
	return fmt.Sprint("invalid Key length ",int(k))
}
type MissingError string
func (m MissingError) Error() string {
	return "Missing: "+string(m)
}

type KeyAcceptor interface{
	// Accepts a public key from a server or a client
	AcceptPublicKey(pk []byte) bool
}
type TrueKeyAcceptor struct{}
func (t TrueKeyAcceptor)AcceptPublicKey(pk []byte) bool{
	return true
}

type ClientConfig struct{
	Acceptor         KeyAcceptor
	ClientPublicKey  []byte
	ClientPrivateKey []byte
	Random           io.Reader
}
func (conf *ClientConfig) Verify() error{
	if e:=verifyLength(conf.ClientPublicKey ,32); e!=nil { return e }
	if e:=verifyLength(conf.ClientPrivateKey,32); e!=nil { return e }
	if conf.Random==nil { return MissingError("*ClientConfig.Random") }
	return nil
}

type ServerConfig struct{
	Acceptor         KeyAcceptor
	ServerPublicKey  []byte
	ServerPrivateKey []byte
	Random           io.Reader
}
func (conf *ServerConfig) Verify() error{
	if e:=verifyLength(conf.ServerPublicKey ,32); e!=nil { return e }
	if e:=verifyLength(conf.ServerPrivateKey,32); e!=nil { return e }
	if conf.Random==nil { return MissingError("*ClientConfig.Random") }
	return nil
}

