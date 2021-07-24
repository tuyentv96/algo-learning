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
			nums: []int{3,4,5,1,2},
			want: 1,
		},
		{
			nums: []int{4,5,6,7,0,1,2},
			want: 0,
		},
	}

	for _, tc := range testcases {
		got := findMin(tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
