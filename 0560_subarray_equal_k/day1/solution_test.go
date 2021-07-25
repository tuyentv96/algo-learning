package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 2, 3},
			target: 3,
			want:   2,
		},
	}

	for _, tc := range testcases {
		got := subarraySum(tc.nums, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
