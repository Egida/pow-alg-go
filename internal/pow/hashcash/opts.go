package hashcash

import (
	"time"
)

type Opt func(*Hashcash)

func WithBits(b int64) Opt {
	return func(hc *Hashcash) {
		hc.bits = b
	}
}

func WithNow(fn func() time.Time) Opt {
	return func(hc *Hashcash) {
		hc.now = fn
	}
}

func WithAttempts(a int64) Opt {
	return func(hc *Hashcash) {
		hc.attempts = a
	}
}
