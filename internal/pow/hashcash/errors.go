package hashcash

import (
	"errors"
)

var (
	ErrSolutionFailed      = errors.New("exceeded limit of attempts to find solution")
	ErrNotAcceptableHeader = errors.New("header is not acceptable")
	ErrNotValidDate        = errors.New("not allowed value for date")

	ErrInvalidHeader         = errors.New("invalid header")
	ErrInvalidHeaderVersion  = errors.New("invalid header version")
	ErrInvalidHeaderBits     = errors.New("invalid header bits")
	ErrInvalidHeaderCounter  = errors.New("invalid header counter")
	ErrInvalidHeaderDate     = errors.New("invalid header date")
	ErrInvalidHeaderResource = errors.New("invalid header resource")
	ErrInvalidHeaderRand     = errors.New("invalid header rand")
)
