package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		want []int
	}{
		{
			arr: []int{3, 5, 1, 6, 2, 0, 8, 0, 0, 7, 4},
		},
	}

	for _, tc := range testcases {
		node := arrayToTree(nil, tc.arr, 0)
		got := subtreeWithAllDeepest(node)
		assert.Equal(t, tc.want, got)
	}
}

func arrayToTree(node *TreeNode, arr []int, index int) *TreeNode {
	if node == nil {
		node = &TreeNode{}
	}

	if index >= len(arr) {
		return nil
	}

	node.Val = arr[index]
	node.Left = arrayToTree(node, arr, 2*index+1)
	node.Right = arrayToTree(node, arr, 2*index+2)

	return node
}
