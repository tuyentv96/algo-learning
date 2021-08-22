package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  [][]int
		want [][]int
	}{
		{
			arr:  [][]int{{1, 0, 2, 1}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want: [][]int{{0, 0, 0, 0}, {3, 0, 5, 2}, {1, 0, 1, 5}},
		},
		{
			arr:  [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tc := range testcases {
		setZeroes(tc.arr)
		assert.Equal(t, tc.want, tc.arr)
	}
}
