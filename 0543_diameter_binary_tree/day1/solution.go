package day1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := height(root.Left), height(root.Right)
	ld, rd := diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)

	return max(1+lh+rh, max(ld, rd))
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
