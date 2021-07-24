package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	ma := New(3)

	testcases := []struct {
		n    int
		want float64
	}{
		{
			n:    1,
			want: 1,
		},
		{
			n:    10,
			want: 10 / 2,
		},
		{
			n:    3,
			want: 14 / 3,
		},
		{
			n:    5,
			want: 6,
		},
	}

	for _, tc := range testcases {
		got := ma.Next(tc.n)
		assert.Equal(t, tc.want, got)
	}
}
