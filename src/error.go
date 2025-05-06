package src

import "errors"

var (
	// connection store
	ErrorDuplicatedSystem = errors.New("system already exists")
	ErrorSystemNotFound   = errors.New("system not found")
)

func AsError[T error](err error) bool {
	var target T
	return errors.As(err, &target)
}
