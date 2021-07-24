package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
	}

	for _, tc := range testcases {
		got := sortedSquares(tc.arr)
		assert.Equal(t, tc.want, got)
	}
}
