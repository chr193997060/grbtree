package grbtree

import "errors"

var (
	errKeyAlreadyExists = errors.New("KeyAlreadyExists")
	errKeyNotExists     = errors.New("KeyNotExists")
	errNotNode          = errors.New("NotNode")
)