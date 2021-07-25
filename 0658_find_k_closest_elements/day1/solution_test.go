package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		x    int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    4,
			x:    9,
			want: []int{1, 2, 3, 4},
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    4,
			x:    -1,
			want: []int{1, 2, 3, 4},
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    4,
			x:    3,
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testcases {
		got := findClosestElements(tc.nums, tc.k, tc.x)
		assert.Equal(t, tc.want, got)
	}
}
