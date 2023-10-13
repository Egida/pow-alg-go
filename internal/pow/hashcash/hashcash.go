package hashcash

import (
	"time"
)

type Hashcash struct {
	bits     int64
	attempts int64
	now      func() time.Time
}

type Opt func(*Hashcash)

const (
	defaultBits     = 20
	defaultAttempts = 1 << 20
)

func New(opts ...Opt) Hashcash {
	hc := Hashcash{
		bits:     defaultBits,
		attempts: defaultAttempts,
		now:      time.Now,
	}

	for _, opt := range opts {
		opt(&hc)
	}

	return hc
}

func (hc Hashcash) Compute(resource string) (Header, error) {
	header := NewHeader(hc.bits, resource, hc.now())

	for !hc.acceptable(header) {
		header.Counter++

		if header.Counter > hc.attempts {
			return header, ErrSolutionFailed
		}
	}

	return header, nil
}

func (hc Hashcash) acceptable(h Header) bool {
	hash := h.Hash()
	for i := range hash[:hc.bits] {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

func (hc Hashcash) Verify(h Header) error {
	if !hc.acceptable(h) {
		return ErrNotAcceptableHeader
	}

	expired := hc.now().AddDate(0, 0, -2)
	if h.Date.After(hc.now()) || h.Date.Before(expired) {
		return ErrNotValidDate
	}

	return nil
}
