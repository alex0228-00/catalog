package src

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSecureCredCodec(t *testing.T) {
	rq := require.New(t)

	mockCred, cleanup := MockCredentials(t, 1)
	defer cleanup()

	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	rq.NoError(err)

	codec, err := NewSecureCredCodec(CipherKey)
	rq.NoError(err)

	text, err := codec.Encode(mockCred[0])
	rq.NoError(err)

	decodedCred, err := codec.Decode(text)
	rq.NoError(err)
	rq.NotNil(decodedCred)
}
