package day1

import (
	"github.com/stretchr/testify/assert"
	"github.com/tuyentv96/algo-learning/structures"
	"testing"
)

func Test(t *testing.T) {
	tcs := []struct {
		root []int
		want []int
	}{
		{
			root: []int{1, 2, 5, 3, 4, structures.NULL, 6},
			want: []int{1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6},
		},
	}

	for _, tc := range tcs {
		root := structures.Ints2TreeNode(tc.root)
		flatten(root)
		assert.Equal(t, tc.want, structures.Tree2Preorder(root))
	}
}
