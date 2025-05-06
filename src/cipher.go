package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

type Cipher struct {
	key    []byte
	cipher cipher.AEAD
}

func NewCipher(key []byte) (*Cipher, error) {
	if len(key) != 32 {
		return nil, errors.Errorf("invalid key length: expected 32 bytes, got %d bytes", len(key))
	}

	block, _ := aes.NewCipher(key)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AES GCM cipher")
	}

	cipher := &Cipher{
		key:    key,
		cipher: aesGCM,
	}
	return cipher, nil
}

func (c *Cipher) Encrypt(data []byte) ([]byte, error) {
	nonce := make([]byte, c.cipher.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.Wrap(err, "failed to generate nonce")
	}

	cipherText := c.cipher.Seal(nonce, nonce, []byte(data), nil)
	return cipherText, nil
}

func (c *Cipher) Decrypt(data []byte) ([]byte, error) {
	nonceSize := c.cipher.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherTextBytes := data[:nonceSize], data[nonceSize:]
	plainText, err := c.cipher.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt data")
	}
	return plainText, nil
}
