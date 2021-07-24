package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{5, 4, 3, 2, 1},
			want: 0,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			nums: []int{2, 3, 4, 5, 1},
			want: 3,
		},
	}

	for _, tc := range testcases {
		got := findPeakElement(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
