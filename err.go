package grbtree

import "errors"

var (
	errKeyAlreadyExists = errors.New("key already exists")
	errKeyNotExists     = errors.New("key no exists")
)