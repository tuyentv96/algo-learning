package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		n    int
		want int
	}{
		{
			arr:  []int{84, -37, 32, 40, 95},
			n:    167,
			want: 3,
		},
	}

	for _, tc := range testcases {
		got := shortestSubarray(tc.arr, tc.n)
		assert.Equal(t, tc.want, got)
	}
}
