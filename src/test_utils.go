package src

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	gomock "go.uber.org/mock/gomock"
)

func MockCredentials(t *testing.T, n int) ([]Credential, func()) {
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
