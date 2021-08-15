package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums [][]int
		want []int
	}{
		{
			nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	for _, tc := range testcases {
		got := spiralOrder(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
