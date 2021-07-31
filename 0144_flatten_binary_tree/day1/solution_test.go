package day1

import (
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: false,
		},
	}

	for _, tc := range testcases {
		flatten(tc.root)
	}
}
