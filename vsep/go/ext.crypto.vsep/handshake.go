package vsep

import "io"
import "crypto/cipher"

func xorEncrypt(dst,add []byte) {
	for i,e := range add {
		dst[i]^=e
	}
}

func ClientHandShake(conf *ClientConfig,r io.Reader,w io.Writer) (io.Reader,io.Writer,error){
	if e:=conf.Verify(); e!=nil { return nil,nil,e }
	srvKey := make([]byte,32)
	_,e := io.ReadFull(r,srvKey)
	if e!=nil { return nil,nil,e }
	if !conf.Acceptor.AcceptPublicKey(srvKey) {
		return nil,nil,KeyAuthenticationFailed{}
	}
	pub ,priv := conf.ClientPublicKey,conf.ClientPrivateKey
	if pub==nil || priv==nil {
		pub,priv,e = CurveKeyPair(conf.Random)
		if e!=nil { return nil,nil,e }
	}
	_,e = w.Write(pub)
	if e!=nil { return nil,nil,e }
	
	c1iv := make([]byte,32)
	c2iv := make([]byte,32)
	
	sessKey := make([]byte,32)
	secr    := make([]byte,32)
	c1kerpu := make([]byte,32)
	c2kerpu := make([]byte,32)
	
	_,e = io.ReadFull(r,sessKey)
	if e!=nil { return nil,nil,e }
	Curve(secr,priv,sessKey)
	_,e = io.ReadFull(r,c1kerpu)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(r,c2kerpu)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(r,c1iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c1kerpu,secr)
	xorEncrypt(c2kerpu,secr)
	xorEncrypt(c1iv,secr)
	
	srvBack,srvSecr,e := CurveSessionKey(conf.Random,srvKey)
	if e!=nil { return nil,nil,e }
	c1kelpu,c1kelpr,e := CurveKeyPair(conf.Random)
	if e!=nil { return nil,nil,e }
	c2kelpu,c2kelpr,e := CurveKeyPair(conf.Random)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(conf.Random,c2iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c1kelpu,srvSecr)
	xorEncrypt(c2kelpu,srvSecr)
	xorEncrypt(c2iv,srvSecr)
	
	_,e = w.Write(srvBack)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c1kelpu)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c2kelpu)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c2iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c2iv,srvSecr) // decrypt it back
	
	c1k := make([]byte,32)
	c2k := make([]byte,32)
	
	Curve(c1k,c1kelpr,c1kerpu)
	Curve(c2k,c2kelpr,c2kerpu)
	
	s2c := CreateDecryptor(c1k,c1iv)
	c2s := CreateEncryptor(c2k,c2iv)
	return  cipher.StreamReader{R:r,S:s2c},
			cipher.StreamWriter{W:w,S:c2s},
			nil
}

func ServerHandShake(conf *ServerConfig,r io.Reader,w io.Writer) (io.Reader,io.Writer,error){
	if e:=conf.Verify(); e!=nil { return nil,nil,e }
	var e error
	pub ,priv := conf.ServerPublicKey,conf.ServerPrivateKey
	if pub==nil || priv==nil {
		pub,priv,e = CurveKeyPair(conf.Random)
		if e!=nil { return nil,nil,e }
	}
	_,e = w.Write(pub)
	if e!=nil { return nil,nil,e }
	
	clntKey := make([]byte,32)
	_,e = io.ReadFull(r,clntKey)
	if e!=nil { return nil,nil,e }
	if !conf.Acceptor.AcceptPublicKey(clntKey) {
		return nil,nil,KeyAuthenticationFailed{}
	}
	
	c1iv := make([]byte,32)
	c2iv := make([]byte,32)
	
	clntBack,clntSecr,e := CurveSessionKey(conf.Random,clntKey)
	if e!=nil { return nil,nil,e }
	c1kelpu,c1kelpr,e := CurveKeyPair(conf.Random)
	if e!=nil { return nil,nil,e }
	c2kelpu,c2kelpr,e := CurveKeyPair(conf.Random)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(conf.Random,c1iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c1kelpu,clntSecr)
	xorEncrypt(c2kelpu,clntSecr)
	xorEncrypt(c1iv,clntSecr)
	
	_,e = w.Write(clntBack)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c1kelpu)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c2kelpu)
	if e!=nil { return nil,nil,e }
	_,e = w.Write(c1iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c1iv,clntSecr) // decrypt it back
	
	sessKey := make([]byte,32)
	secr    := make([]byte,32)
	c1kerpu := make([]byte,32)
	c2kerpu := make([]byte,32)
	
	_,e = io.ReadFull(r,sessKey)
	if e!=nil { return nil,nil,e }
	Curve(secr,priv,sessKey)
	_,e = io.ReadFull(r,c1kerpu)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(r,c2kerpu)
	if e!=nil { return nil,nil,e }
	_,e = io.ReadFull(r,c2iv)
	if e!=nil { return nil,nil,e }
	
	xorEncrypt(c1kerpu,secr)
	xorEncrypt(c2kerpu,secr)
	xorEncrypt(c2iv,secr)
	
	c1k := make([]byte,32)
	c2k := make([]byte,32)
	
	Curve(c1k,c1kelpr,c1kerpu)
	Curve(c2k,c2kelpr,c2kerpu)
	
	s2c := CreateEncryptor(c1k,c1iv)
	c2s := CreateDecryptor(c2k,c2iv)
	return  cipher.StreamReader{R:r,S:c2s},
			cipher.StreamWriter{W:w,S:s2c},
			nil
}



