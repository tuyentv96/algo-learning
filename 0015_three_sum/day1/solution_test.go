package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, tc := range testcases {
		got := threeSum(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
