package day1

import (
	"github.com/stretchr/testify/assert"
	"github.com/tuyentv96/algo-learning/structures"
	"testing"
)

func Test(t *testing.T) {
	tcs := []struct {
		root []int
		want [][]int
	}{
		{
			root: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, tc := range tcs {
		got := levelOrder((*TreeNode)(structures.Ints2TreeNode(tc.root)))
		assert.Equal(t, tc.want, got)
	}
}
