package day1

import "github.com/tuyentv96/algo-learning/structures"

type TreeNode structures.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}

	var prev *TreeNode
	cur := root

	for cur != nil {
		prev = cur

		if val < cur.Val {
			cur = (*TreeNode)(cur.Left)
		} else {
			cur = (*TreeNode)(cur.Right)
		}
	}

	if val < prev.Val {
		prev.Left = (*structures.TreeNode)(newNode)
	} else {
		prev.Right = (*structures.TreeNode)(newNode)
	}

	return root
}
