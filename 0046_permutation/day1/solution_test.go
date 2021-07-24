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
			nums: []int{1, 2, 3},
			want: 6,
		},
	}

	for _, tc := range testcases {
		got := permute(tc.nums)
		assert.Equal(t, tc.want, len(got))
	}
}
