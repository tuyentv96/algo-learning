package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
	}

	for _, tc := range testcases {
		got := findKthLargest(tc.nums, tc.k)
		assert.Equal(t, tc.want, got)
	}
}
