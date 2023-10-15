package redis

import (
	"errors"
)

var ErrAlreadyExists = errors.New("key already exists")
