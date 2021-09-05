package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums [][]int
		n    int
		want bool
	}{
		{
			nums: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			n:    5,
			want: true,
		},
	}

	for _, tc := range testcases {
		got := validTree(tc.n, tc.nums)
		assert.Equal(t, tc.want, got)
	}
}
