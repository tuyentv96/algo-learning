package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		want   bool
	}{
		{
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			target: 13,
			want:   true,
		},
		{
			nums:   []int{1, 0, 1, 1, 1},
			target: 0,
			want:   true,
		},
		{
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
	}

	for _, tc := range testcases {
		got := search(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
