package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr    []int
		target int
		want   [][]int
	}{
		{
			arr:    []int{2, 3, 5},
			target: 8,
			want: [][]int{
				{
					2, 2, 2, 2,
				},
			},
		},
	}

	for _, tc := range testcases {
		got := combinationSum(tc.arr, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
