package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{4, 7, 9, 10},
			k:    3,
			want: 8,
		},
		{
			nums: []int{4, 7, 9, 10},
			k:    1,
			want: 5,
		},
	}

	for _, tc := range testcases {
		got := find(tc.nums, tc.k)
		assert.Equal(t, tc.want, got)
	}
}
