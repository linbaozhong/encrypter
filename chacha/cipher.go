package chacha

import (
	"github.com/aead/chacha20"
)

type ChaCha struct {
	key []byte
	iv  []byte
}

func New(key, iv []byte) *ChaCha {
	return &ChaCha{key: key, iv: iv}
}

func (c *ChaCha) Encode(plainData []byte) (cipherData []byte, err error) {
	stream, err := chacha20.NewCipher(c.iv, c.key)
	if err != nil {
		return
	}
	cipherData = make([]byte, len(plainData))
	stream.XORKeyStream(cipherData, plainData)
	return
}

func (c *ChaCha) Decode(cipherText []byte) (plainText []byte, err error) {
	stream, err := chacha20.NewCipher(c.iv, c.key)
	if err != nil {
		return
	}
	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	return
}
