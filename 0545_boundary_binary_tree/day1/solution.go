package day1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundary(root *TreeNode) []int {
	var result []int
	leftBoundary(root, &result)
	leaf(root, &result)
	rightBoundary(root, &result)
	return result
}

func leftBoundary(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	if root.Left != nil {
		leftBoundary(root.Left, result)
	} else if root.Right != nil {
		leftBoundary(root.Left, result)
	}
}

func rightBoundary(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	if root.Right != nil {
		leftBoundary(root.Right, result)
	} else if root.Left != nil {
		leftBoundary(root.Right, result)
	}
}

func leaf(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	leaf(root.Left, result)
	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
		return
	}
	leaf(root.Right, result)
}
