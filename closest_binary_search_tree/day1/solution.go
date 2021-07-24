package day1

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestBinarySearchTree(root *TreeNode, k int) int {
	minDiff := 1 << 32
	helper(root, k, &minDiff)
	return minDiff
}

func helper(root *TreeNode, k int, minDiff *int) {
	if root == nil {
		return
	}

	diff := int(math.Abs(float64(root.Val - k)))
	if diff < *minDiff {
		minDiff = &diff
	}

	if root.Val < k {
		helper(root.Left, k, minDiff)
	} else {
		helper(root.Right, k, minDiff)
	}
}
