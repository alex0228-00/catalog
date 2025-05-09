//go:generate mockgen -source=cred.go -destination=mock_cred.go -package=service

package service

import (
	"encoding/base64"

	"catalog/src"

	"github.com/pkg/errors"
)

type CredentialType uint8
type credConstructor func() Credential

const (
	CredTypeToken CredentialType = iota
)

type Credential interface {
	Type() CredentialType
	Encode() ([]byte, error)
	Decode(data []byte) error
}

type SecureCredCodec struct {
	cipher *src.Cipher
	creds  map[CredentialType]credConstructor
}

func NewSecureCredCodec(key []byte) (*SecureCredCodec, error) {
	cipher, err := src.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cipher")
	}
	return &SecureCredCodec{cipher: cipher, creds: creds}, nil
}

func (codec *SecureCredCodec) Encode(cred Credential) (string, error) {
	encoded, err := cred.Encode()
	if err != nil {
		return "", errors.Wrap(err, "failed to encode credential")
	}

	raw := src.EncodeAsTLVBlock(byte(cred.Type()), encoded)

	encrypted, err := codec.cipher.Encrypt(raw)
	if err != nil {
		return "", errors.Wrap(err, "failed to encrypt credential")
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (codec *SecureCredCodec) Decode(text string) (Credential, error) {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode base64 credential")
	}

	decrypted, err := codec.cipher.Decrypt(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt credential")
	}

	_type, rawCred, err := src.DecodeFromTLVBlock(decrypted)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode TLV block")
	}

	constructor, ok := codec.creds[CredentialType(_type)]
	if !ok {
		return nil, errors.Errorf("unsupported credential type: %d", _type)
	}
	cred := constructor()

	if err := cred.Decode(rawCred); err != nil {
		return nil, errors.Wrap(err, "failed to decode credential data")
	}
	return cred, nil
}

var creds = map[CredentialType]credConstructor{
	CredTypeToken: func() Credential { return &TokenCredential{} },
}

type TokenCredential struct {
	token []byte
}

func (t *TokenCredential) Type() CredentialType {
	return CredTypeToken
}

func (t *TokenCredential) Decode(data []byte) error {
	t.token = make([]byte, len(data))
	copy(t.token, data)
	return nil
}

func (t *TokenCredential) Encode() ([]byte, error) {
	token := make([]byte, len(t.token))
	copy(token, t.token)
	return token, nil
}
