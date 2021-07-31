package day1

import (
	"fmt"
	"github.com/tuyentv96/algo-learning/structures"
	"testing"
)

func Test_Problem114(t *testing.T) {

	qs := []question114{

		{
			para114{[]int{1, 2, 3, 4, 5, 6}},
			ans114{[]int{1, 2, 3, 4, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 114------------------------\n")

	for _, q := range qs {
		_, p := q.ans114, q.para114
		fmt.Printf("【input】:%v       \n", p)
		rootOne := structures.Ints2TreeNode(p.one)
		flatten(rootOne)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(rootOne))
	}
	fmt.Printf("\n\n\n")
}
