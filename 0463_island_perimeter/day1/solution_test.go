package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  [][]int
		want int
	}{
		{
			arr: [][]int{
				{0, 1},
			},
			want: 4,
		},
	}

	for _, tc := range testcases {
		got := islandPerimeter(tc.arr)
		assert.Equal(t, tc.want, got)
	}
}
