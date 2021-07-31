package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{3, 0, 1, 0},
			k:    1,
			want: []int{0},
		},
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
	}

	for _, tc := range testcases {
		got := topKFrequent(tc.nums, tc.k)
		assert.Equal(t, tc.want, got)
	}
}
