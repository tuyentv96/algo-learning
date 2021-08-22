package day1

import "github.com/tuyentv96/algo-learning/structures"

type TreeNode structures.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree((*TreeNode)(root.Right))
	invertTree((*TreeNode)(root.Left))

	return root
}
