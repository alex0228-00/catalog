package src

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCipher(t *testing.T) {
	rq := require.New(t)

	expected := "Hello, World!"
	cipher, err := NewCipher(CipherKey)
	rq.NoError(err)

	cipherText, err := cipher.Encrypt([]byte(expected))
	rq.NoError(err)

	text, err := cipher.Decrypt(cipherText)
	rq.NoError(err)
	rq.Equal(expected, string(text))
}
