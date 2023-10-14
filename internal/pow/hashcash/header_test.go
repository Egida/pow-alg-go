package hashcash_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ivasilkov/pow-alg-go/internal/pow/hashcash"
)

func TestParseHeader(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		testCases := []struct {
			name string
			in   string
			err  error
		}{
			{
				name: "invalid_header",
				in:   "invalid:format",
				err:  hashcash.ErrInvalidHeader,
			},
			{
				name: "invalid_header_format",
				in:   "1:20:1697280348:any-client-data:asd:KN5rx3yHb6M=:380479",
				err:  hashcash.ErrInvalidHeader,
			},
			{
				name: "invalid_header_version",
				in:   "asd:20:1697280348:any-client-data::KN5rx3yHb6M=:380479",
				err:  hashcash.ErrInvalidHeaderVersion,
			},
			{
				name: "invalid_header_bits",
				in:   "1:asd:1697280348:any-client-data::KN5rx3yHb6M=:380479",
				err:  hashcash.ErrInvalidHeaderBits,
			},
			{
				name: "invalid_header_counter",
				in:   "1:20:1697280348:any-client-data::KN5rx3yHb6M=:asd",
				err:  hashcash.ErrInvalidHeaderCounter,
			},
			{
				name: "invalid_header_date",
				in:   "1:20:asd:any-client-data::KN5rx3yHb6M=:380479",
				err:  hashcash.ErrInvalidHeaderDate,
			},
			{
				name: "invalid_header_resource",
				in:   "1:20:1697280348:::KN5rx3yHb6M=:380479",
				err:  hashcash.ErrInvalidHeaderResource,
			},
			{
				name: "invalid_header_rand",
				in:   "1:20:1697280348:any-client-data:::380479",
				err:  hashcash.ErrInvalidHeaderRand,
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				_, err := hashcash.ParseHeader(tc.in)
				require.ErrorIs(t, err, tc.err)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		in := "1:20:1697280348:any-client-data::KN5rx3yHb6M=:380479"

		h, err := hashcash.ParseHeader(in)

		require.NoError(t, err)
		require.Equal(t, int64(1), h.Ver)
		require.Equal(t, int64(20), h.Bits)
		require.Equal(t, time.Unix(1697280348, 0), h.Date)
		require.Equal(t, "any-client-data", h.Resource)
		require.Equal(t, "KN5rx3yHb6M=", h.Rand)
		require.Equal(t, int64(380479), h.Counter)
	})
}
