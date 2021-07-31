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
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
	}

	for _, tc := range testcases {
		got := lengthOfLIS(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
