package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
	}

	for _, tc := range testcases {
		got := search(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
