package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  string
		want *TreeNode
	}{
		{
			str: "4(2(3)(1))(6(5))",
		},
	}

	for _, tc := range testcases {
		got := construct(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
