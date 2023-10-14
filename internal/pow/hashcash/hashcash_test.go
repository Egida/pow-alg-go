package hashcash_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
)

func TestHashcash_Compute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		hc := hashcash.New(hashcash.WithBits(8))
		h, err := hc.Compute("resource")

		require.NoError(t, err)
		require.Equal(t, "00", h.Hash()[:2])
	})

	t.Run("error", func(t *testing.T) {
		hc := hashcash.New(hashcash.WithAttempts(1))

		h, err := hc.Compute("resource")

		require.ErrorIs(t, err, hashcash.ErrSolutionFailed)
		require.Equal(t, int64(2), h.Counter)
	})
}

func TestHashcash_Verify(t *testing.T) {
	hTime := time.Date(2023, 10, 14, 11, 3, 47, 0, time.UTC)
	h := hashcash.Header{
		Ver:      1,
		Bits:     20,
		Counter:  950363,
		Date:     hTime,
		Resource: "resource",
		Rand:     "rwPoU2oqz3I=",
	}

	t.Run("success", func(t *testing.T) {
		hc := hashcash.New(
			hashcash.WithNow(func() time.Time {
				return hTime.Add(time.Hour)
			}),
		)

		err := hc.Verify(h)

		require.NoError(t, err)
	})

	t.Run("not_acceptable", func(t *testing.T) {
		hc := hashcash.New()

		in := hashcash.Header{}
		err := hc.Verify(in)

		require.ErrorIs(t, err, hashcash.ErrNotAcceptableHeader)
	})

	t.Run("future_date", func(t *testing.T) {
		hc := hashcash.New(
			hashcash.WithNow(func() time.Time {
				return hTime.Add(-time.Hour)
			}),
		)

		err := hc.Verify(h)

		require.ErrorIs(t, err, hashcash.ErrNotValidDate)
	})

	t.Run("expired_date", func(t *testing.T) {
		hc := hashcash.New(
			hashcash.WithNow(func() time.Time {
				return hTime.Add(time.Hour * 72)
			}),
		)

		err := hc.Verify(h)

		require.ErrorIs(t, err, hashcash.ErrNotValidDate)
	})
}
