package day1

import "github.com/tuyentv96/algo-learning/structures"

type TreeNode structures.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode

	for true {
		for root != nil {
			stack = append(stack, root)
			root = (*TreeNode)(root.Left)
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = (*TreeNode)(root.Right)
	}

	return 0
}
