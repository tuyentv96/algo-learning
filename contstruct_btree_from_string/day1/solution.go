package day1

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func construct(str string) *TreeNode {
	if str == "" {
		return nil
	}

	root, _ := helper(str, 0)
	return root
}

// "4(2(3)(1))(6(5))"
func helper(str string, i int) (*TreeNode, int) {
	p := i
	root := &TreeNode{}
	for p < len(str) && str[p] != '(' && str[p] != ')' {
		p++
	}

	// get root
	tmp := str[i:p]
	val, _ := strconv.Atoi(tmp)
	root.Val = val

	if p < len(str) && str[p] == '(' {
		p++
		root.Left, p = helper(str, p)
	}

	if p < len(str) && str[p] == ')' {
		p++
		return root, p
	}

	if p < len(str) && str[p] == '(' {
		p++
		root.Right, p = helper(str, p)
	}

	if p < len(str) && str[p] == ')' {
		p++
		return root, p
	}

	return root, p
}
