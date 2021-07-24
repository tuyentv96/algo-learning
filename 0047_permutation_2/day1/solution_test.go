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
			nums: []int{1, 1, 2},
			want: 3,
		},
	}

	for _, tc := range testcases {
		got := permuteUnique(tc.nums)
		assert.Equal(t, tc.want, len(got))
	}
}
