package day1

import (
	"github.com/stretchr/testify/assert"
	"github.com/tuyentv96/algo-learning/structures"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{2, 1, 3},
			want: true,
		},
	}

	for _, tc := range testcases {
		got := isValidBST((*TreeNode)(structures.Ints2TreeNode(tc.nums)))
		assert.Equal(t, tc.want, got)
	}
}
