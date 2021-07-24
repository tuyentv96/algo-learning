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
			arr:  []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
	}

	for _, tc := range testcases {
		sortColors(tc.arr)
		assert.Equal(t, tc.want, tc.arr)
	}
}
