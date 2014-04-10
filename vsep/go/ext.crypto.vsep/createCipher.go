package vsep

import "crypto/aes"
import "crypto/cipher"
import "ext.crypto.vsep/ecfb"

func newAesCipher(key []byte) cipher.Block {
	c,e := aes.NewCipher(key)
	if e!=nil { panic(e) }
	return c
}

func CreateEncryptor(key []byte, iv []byte) cipher.Stream{
	bc1,iv1 := newAesCipher(key[:16]),iv[:16] // 1..16
	bc2,iv2 := newAesCipher(key[16:]),iv[:16] // 16..32
	return ecfb.NewECFBEncrypter(bc2,cipher.NewCTR(bc1,iv1),iv2)
}
func CreateDecryptor(key []byte, iv []byte) cipher.Stream{
	bc1,iv1 := newAesCipher(key[:16]),iv[:16] // 1..16
	bc2,iv2 := newAesCipher(key[16:]),iv[:16] // 16..32
	return ecfb.NewECFBDecrypter(bc2,cipher.NewCTR(bc1,iv1),iv2)
}

