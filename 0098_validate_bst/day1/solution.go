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

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	return dfs(root, &prev)
}

func dfs(root *TreeNode, prev **TreeNode) bool {
	if root == nil {
		return true
	}

	if !dfs((*TreeNode)(root.Left), prev) {
		return false
	}

	if *prev != nil && root.Val <= (*prev).Val {
		return false
	}
	*prev = root

	return dfs((*TreeNode)(root.Right), prev)
}

//func isValidBST(root *TreeNode) bool {
//	return dfs(root, nil, nil)
//}
//
//func dfs(root *TreeNode, low, high *int) bool {
//	if root == nil {
//		return true
//	}
//
//	if (low != nil && root.Val <= *low) || (high != nil && root.Val >= *high) {
//		return false
//	}
//
//	return dfs((*TreeNode)(root.Left), low, &root.Val) && dfs((*TreeNode)(root.Right), &root.Val, high)
//}
