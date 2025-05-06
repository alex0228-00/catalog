package src

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTLVEncodeDecode(t *testing.T) {
	rq := require.New(t)
	text := "Hello, World!"
	_type := byte(1)

	block := EncodeAsTLVBlock(_type, []byte(text))
	decodedType, decodedText, err := DecodeFromTLVBlock(block)
	rq.NoError(err)
	rq.Equal(_type, decodedType)
	rq.Equal(text, string(decodedText))
}
