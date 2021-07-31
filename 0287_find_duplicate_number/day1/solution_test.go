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
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
	}

	for _, tc := range testcases {
		got := findDuplicate(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
