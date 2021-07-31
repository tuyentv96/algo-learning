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
			nums: []int{1, 2, 3, 4, 6, 7, 8},
			want: 5,
		},
	}

	for _, tc := range testcases {
		got := find(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
