package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums   [][]int
		target int
		want   bool
	}{
		{
			nums:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			nums:   [][]int{{1, 1}},
			target: 2,
			want:   false,
		},
	}

	for _, tc := range testcases {
		got := searchMatrix(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
