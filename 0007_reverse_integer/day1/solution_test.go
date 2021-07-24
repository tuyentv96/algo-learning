package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		x    int
		want int
	}{
		{
			x:    -123,
			want: -321,
		},
		{
			x:    123,
			want: 321,
		},
	}

	for _, tc := range testcases {
		got := reverse(tc.x)
		assert.Equal(t, tc.want, got)
	}
}
