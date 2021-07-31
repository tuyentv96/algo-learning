package day1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {

}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	right := root.Right
	root.Right = root.Left
	traverse(root.Left)
	traverse(right)
}
