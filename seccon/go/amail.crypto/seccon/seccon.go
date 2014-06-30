package seccon

import "io"
import "crypto/aes"
import "crypto/cipher"
import "amail.crypto/ecfb"
import "amail.crypto/curve25519"

func newAesCipher(key []byte) cipher.Block {
	c,e := aes.NewCipher(key)
	if e!=nil { panic(e) }
	return c
}

func createEncryptor(key []byte) cipher.Stream{
	// use the key as IV!
	bc1,iv1 := newAesCipher(key[:16]),key[:16] // 1..16
	bc2,iv2 := newAesCipher(key[16:]),key[:16] // 16..32
	return ecfb.NewECFBEncrypter(bc2,cipher.NewCTR(bc1,iv1),iv2)
}
func createDecryptor(key []byte) cipher.Stream{
	// use the key as IV!
	bc1,iv1 := newAesCipher(key[:16]),key[:16] // 1..16
	bc2,iv2 := newAesCipher(key[16:]),key[:16] // 16..32
	return ecfb.NewECFBDecrypter(bc2,cipher.NewCTR(bc1,iv1),iv2)
}

func curve(out, in, base []byte){
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
func curveSessionKey(rand io.Reader,pub []byte) ([]byte,[]byte,error){
	t := make([]byte,32)
	T := make([]byte,32)
	K := make([]byte,32)
	_,e := io.ReadFull(rand,t)
	if e!=nil { return nil,nil,e }
	t[ 0] &= 248
	t[31] &= 127
	t[31] |= 64
	curve(T,t,nil)
	curve(K,t,pub)
	return T,K,nil
}

// utility function. Creates an curve25519 Key Pair
// returns PublicKey, PrivateKey, Error
func GenerateKeyPair(rand io.Reader) ([]byte,[]byte,error){
	t := make([]byte,32)
	T := make([]byte,32)
	_,e := io.ReadFull(rand,t)
	if e!=nil { return nil,nil,e }
	t[ 0] &= 248
	t[31] &= 127
	t[31] |= 64
	curve(T,t,nil)
	return T,t,nil
}

func EncryptSeccon(rand io.Reader,pub []byte, dst io.Writer) (io.Writer,error) {
	T,K,e := curveSessionKey(rand,pub)
	if e!=nil { return nil,e }
	n,e := dst.Write(T)
	if n<32 && e!=nil { return nil,e }
	ciph := createEncryptor(K)
	return cipher.StreamWriter{W:dst,S:ciph},nil
}
func DecryptSeccon(priv []byte, src io.Reader) (io.Reader,error) {
	T := make([]byte,32)
	K := make([]byte,32)
	n,e := io.ReadFull(src,T)
	if n<32 && e!=nil { return nil,e }
	curve(K,priv,T)
	ciph := createDecryptor(K)
	return cipher.StreamReader{R:src,S:ciph},nil
}


