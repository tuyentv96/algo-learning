package day1

import "github.com/tuyentv96/algo-learning/structures"

type TreeNode structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := leftDepth((*TreeNode)(root.Left))
	rd := leftDepth((*TreeNode)(root.Right))
	if ld == rd {
		return 1<<ld + countNodes((*TreeNode)(root.Right))
	}

	return 1<<rd + countNodes((*TreeNode)(root.Left))
}

func leftDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + leftDepth((*TreeNode)(root.Left))
}
