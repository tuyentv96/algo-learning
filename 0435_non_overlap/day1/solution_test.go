package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums [][]int
		want int
	}{
		{
			nums: [][]int{{1, 2}, {2, 3}},
			want: 0,
		},
		{
			nums: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want: 1,
		},
	}

	for _, tc := range testcases {
		got := eraseOverlapIntervals(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
