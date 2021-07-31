package day1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeToDoubleLinkedList(root *TreeNode) *TreeNode {
	var head **TreeNode
	var prev **TreeNode
	dfs(root, head, prev)
	return *head
}

func dfs(root *TreeNode, head **TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left, head, prev)
	if prev == nil {
		*head = root
	} else {
		(*prev).Right = root
		root.Left = (*prev)
	}
	*prev = root
	dfs(root.Right, head, prev)
}
