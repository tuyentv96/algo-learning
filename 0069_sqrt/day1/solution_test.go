package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums int
		want int
	}{
		{
			nums: 8,
			want: 2,
		},
		{
			nums: 4,
			want: 2,
		},
	}

	for _, tc := range testcases {
		got := mySqrt(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
