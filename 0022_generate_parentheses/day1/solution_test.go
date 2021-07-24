package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		n    int
		want int
	}{
		{
			n:    2,
			want: 2,
		},
	}

	for _, tc := range testcases {
		got := generateParenthesis(tc.n)
		assert.Equal(t, tc.want, got)
	}
}
