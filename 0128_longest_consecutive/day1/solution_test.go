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
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
	}

	for _, tc := range testcases {
		got := longestConsecutive(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
