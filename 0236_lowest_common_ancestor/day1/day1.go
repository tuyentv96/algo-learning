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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor((*TreeNode)(root.Left), p, q)
	right := lowestCommonAncestor((*TreeNode)(root.Right), p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
