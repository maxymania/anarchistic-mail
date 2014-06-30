// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package ecfb

import "crypto/cipher"

func dup(p []byte) []byte {
	q := make([]byte, len(p))
	copy(q, p)
	return q
}

type cfb struct {
	b       cipher.Block
	encs    cipher.Stream
	out     []byte
	outUsed int
	decrypt bool
}

// NewECFBEncrypter returns a Stream which encrypts with Encrypted cipher feedback mode,
// using the given Block. The iv must be the same length as the Block's block
// size.
func NewECFBEncrypter(block cipher.Block, strm cipher.Stream, iv []byte) cipher.Stream {
	if len(iv) != block.BlockSize() {
		panic("ecfb.NewECBFEncrypter: IV length must equal block size")
	}
	return newECFB(block, strm, iv, false)
}

// NewECFBDecrypter returns a Stream which decrypts with Encrypted cipher feedback mode,
// using the given Block. The iv must be the same length as the Block's block
// size.
func NewECFBDecrypter(block cipher.Block, strm cipher.Stream, iv []byte) cipher.Stream {
	if len(iv) != block.BlockSize() {
		panic("cipher.NewCBFEncrypter: IV length must equal block size")
	}
	return newECFB(block, strm, iv, true)
}

func newECFB(block cipher.Block, strm cipher.Stream, iv []byte, decrypt bool) cipher.Stream {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		return nil
	}

	x := &cfb{
		b:       block,
		encs:    strm,
		out:     make([]byte, blockSize),
		outUsed: 0,
		decrypt: decrypt,
	}
	block.Encrypt(x.out, iv)

	return x
}

func (x *cfb) XORKeyStream(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		if x.outUsed == len(x.out) {
			x.encs.XORKeyStream(x.out, x.out)
			x.b.Encrypt(x.out, x.out)
			x.outUsed = 0
		}

		if x.decrypt {
			t := src[i]
			dst[i] = src[i] ^ x.out[x.outUsed]
			x.out[x.outUsed] = t
		} else {
			x.out[x.outUsed] ^= src[i]
			dst[i] = x.out[x.outUsed]
		}
		x.outUsed++
	}
}
