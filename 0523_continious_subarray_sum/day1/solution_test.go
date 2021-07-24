package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		n    int
		want bool
	}{
		{
			arr:  []int{3, 3},
			n:    6,
			want: true,
		},
		{
			arr:  []int{23, 2, 4, 6, 6},
			n:    7,
			want: true,
		},
	}

	for _, tc := range testcases {
		got := checkSubarraySum(tc.arr, tc.n)
		assert.Equal(t, tc.want, got)
	}
}
