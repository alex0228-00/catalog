package src

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Env Constants
const (
	EnvServerHost = "CATALOG_SERVER_HOST"
	EnvServerPort = "CATALOG_SERVER_PORT"

	EnvCipherKey = "CATALOG_CIPHER_KEY"

	EnvDbHost            = "CATALOG_DB_HOST"
	EnvDbPort            = "CATALOG_DB_PORT"
	EnvDbRootUser        = "CATALOG_DB_ROOT_USER"
	EnvDbRootPassword    = "CATALOG_DB_ROOT_PASSWORD"
	EnvDbUser            = "CATALOG_DB_USER"
	EnvDbPassword        = "CATALOG_DB_PASSWORD"
	EnvDbSupportUser     = "CATALOG_DB_SUPPORT_USER"
	EnvDbSupportPassword = "CATALOG_DB_SUPPORT_PASSWORD"
	EnvDbSchema          = "CATALOG_DB_SCHEMA"
)

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return value
}

// TLV encoder/decoder
// +------+-----+-------+
// | TYPE | LEN | VALUE |
// +------+-----+-------+
// |   1  |  8  |  LEN  |
// +------+-----+-------+

func EncodeAsTLVBlock(_type byte, data []byte) []byte {
	raw := make([]byte, 1+8+len(data))
	raw[0] = _type
	binary.BigEndian.PutUint64(raw[1:], uint64(len(data)))
	copy(raw[9:], data)
	return raw
}

func DecodeFromTLVBlock(data []byte) (byte, []byte, error) {
	if len(data) < 9 {
		return 0, nil, errors.Errorf(
			"invalid TLV block length: expected at least 9 bytes, got %d bytes", len(data),
		)
	}

	length := binary.BigEndian.Uint64(data[1:9])
	if len(data[9:]) != int(length) {
		return 0, nil, errors.Errorf(
			"invalid TLV block length: expected at least 9 bytes, got %d bytes", len(data),
		)
	}

	return data[0], data[9 : 9+length], nil
}

// Only for tests
var CipherKey = []byte("12345678901234567890123456789012")

// DB related
type Dsn struct {
	Type     string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func (d Dsn) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", d.Username, d.Password, d.Host, d.Port, d.Database)
}

// MultipleError 用来收集多个错误
type MultipleError struct {
	message string
	Errors  []error
}

func NewMultipleError(message string) *MultipleError {
	return &MultipleError{
		message: message,
	}
}

func (m *MultipleError) Error() string {
	var msgs []string
	for _, err := range m.Errors {
		msgs = append(msgs, err.Error())
	}
	return fmt.Sprintf("%s: %s", m.message, strings.Join(msgs, "; "))
}

// 如果没有错误，就返回 nil
func (m *MultipleError) ErrOrNil() error {
	if len(m.Errors) == 0 {
		return nil
	}
	return m
}

func (m *MultipleError) Add(err error) {
	m.Errors = append(m.Errors, err)
}
