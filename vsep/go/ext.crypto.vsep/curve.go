package vsep

import "io"
import "ext.crypto.vsep/curve25519"

func Curve(out, in, base []byte){
	var o,i,b [32]byte
	copy(i[:],in)
	if base!=nil {
		copy(b[:],base)
		curve25519.ScalarMult(&o,&i,&b)
	} else {
		curve25519.ScalarBaseMult(&o,&i)
	}
	copy(out,o[:])
}
func CurveSessionKey(rand io.Reader,pub []byte) ([]byte,[]byte,error){
	t := make([]byte,32)
	T := make([]byte,32)
	K := make([]byte,32)
	_,e := io.ReadFull(rand,t)
	if e!=nil { return nil,nil,e }
	t[ 0] &= 248
	t[31] &= 127
	t[31] |= 64
	Curve(T,t,nil)
	Curve(K,t,pub)
	return T,K,nil
}
func CurveKeyPair(rand io.Reader) ([]byte,[]byte,error){
	t := make([]byte,32)
	T := make([]byte,32)
	_,e := io.ReadFull(rand,t)
	if e!=nil { return nil,nil,e }
	t[ 0] &= 248
	t[31] &= 127
	t[31] |= 64
	Curve(T,t,nil)
	return T,t,nil
}

