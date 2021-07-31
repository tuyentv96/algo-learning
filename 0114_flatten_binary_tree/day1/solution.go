package day1

import "github.com/tuyentv96/algo-learning/structures"

type TreeNode = structures.TreeNode

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func flatten(root *TreeNode) {
	var prev *TreeNode
	dfs(root, &prev)
}

func dfs(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Right, prev)
	dfs(root.Left, prev)
	root.Right = *prev
	root.Left = nil
	*prev = root
}
