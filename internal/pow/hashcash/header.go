package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	Ver      int64
	Bits     int64
	Counter  int64
	Date     time.Time
	Resource string
	Rand     string
}

func NewHeader(bits int64, resource string, now time.Time) Header {
	token := make([]byte, 8)
	rand.Read(token)

	return Header{
		Ver:      1,
		Bits:     bits,
		Date:     now,
		Resource: resource,
		Rand:     base64.StdEncoding.EncodeToString(token),
		Counter:  0,
	}
}

func (h *Header) String() string {
	return fmt.Sprintf("%d:%d:%d:%s::%s:%d", h.Ver, h.Bits, h.Date.Unix(), h.Resource, h.Rand, h.Counter)
}

func (h *Header) Hash() string {
	hash := sha1.New()
	hash.Write([]byte(h.String()))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

const hashcashLength = 7

func ParseHeader(header string) (Header, error) {
	var out Header

	vals := strings.Split(header, ":")
	if len(vals) != hashcashLength || vals[4] != "" {
		return out, ErrInvalidHeader
	}

	var err error
	out.Ver, err = strconv.ParseInt(vals[0], 10, 64)
	if err != nil {
		return out, ErrInvalidHeaderVersion
	}

	out.Bits, err = strconv.ParseInt(vals[1], 10, 64)
	if err != nil {
		return out, ErrInvalidHeaderBits
	}

	out.Counter, err = strconv.ParseInt(vals[6], 10, 64)
	if err != nil {
		return out, ErrInvalidHeaderCounter
	}

	unix, err := strconv.ParseInt(vals[2], 10, 64)
	if err != nil {
		return out, ErrInvalidHeaderDate
	}
	out.Date = time.Unix(unix, 0)

	out.Resource = vals[3]
	if out.Resource == "" {
		return out, ErrInvalidHeaderResource
	}

	out.Rand = vals[5]
	if out.Resource == "" {
		return out, ErrInvalidHeaderRand
	}

	return out, nil
}
