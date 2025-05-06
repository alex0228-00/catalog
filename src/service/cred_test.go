package service

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"catalog/src"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestSecureCredCodec(t *testing.T) {
	rq := require.New(t)

	mockCred, cleanup := mockCredentials(t, 1)
	defer cleanup()

	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	rq.NoError(err)

	codec, err := NewSecureCredCodec(src.CipherKey)
	rq.NoError(err)

	text, err := codec.Encode(mockCred[0])
	rq.NoError(err)

	decodedCred, err := codec.Decode(text)
	rq.NoError(err)
	rq.NotNil(decodedCred)
}

func mockCredentials(t *testing.T, n int) ([]Credential, func()) {
	ctrl := gomock.NewController(t)
	mockCreds := make(map[CredentialType]credConstructor, n)
	ret := make([]Credential, 0, n)

	for i := range n {
		text := fmt.Sprintf("mocked data %d", i)
		_type := CredentialType(i)

		mockCred := NewMockCredential(ctrl)
		mockCred.EXPECT().Type().Return(_type).AnyTimes()
		mockCred.EXPECT().Encode().Return([]byte(text), nil).AnyTimes()
		mockCred.EXPECT().Decode(gomock.Any()).DoAndReturn(func(data []byte) error {
			if string(data) != text {
				return errors.New("data mismatch")
			}
			return nil
		})

		mockCreds[_type] = func() Credential { return mockCred }
		ret = append(ret, mockCred)
	}

	oldCreds := creds
	creds = mockCreds

	return ret, func() {
		creds = oldCreds
		ctrl.Finish()
	}
}
