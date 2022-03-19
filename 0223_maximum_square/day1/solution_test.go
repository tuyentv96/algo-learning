package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test223(t *testing.T) {

	tcs := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
	}

	for _, tc := range tcs {
		got := maximalSquare(tc.matrix)
		assert.Equalf(t, tc.want, got, "want: %d, got: %d", tc.want, got)
	}
}
