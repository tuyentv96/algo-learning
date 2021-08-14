package day1

import (
	"github.com/stretchr/testify/assert"
	"github.com/tuyentv96/algo-learning/structures"
	"testing"
)

func Test(t *testing.T) {
	tcs := []struct {
		root []int
		want []int
	}{
		{
			root: []int{-10, -3, 0, 5, 9},
			want: []int{0, -3, 9, -10, structures.NULL, 5},
		},
	}

	for _, tc := range tcs {
		got := sortedListToBST(structures.Ints2List(tc.root))
		arr := []int{}
		structures.T2s(got, &arr)
		assert.Equal(t, tc.want, arr)
	}
}
