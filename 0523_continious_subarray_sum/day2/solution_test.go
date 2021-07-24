package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		n    int
		want bool
	}{
		{
			arr:  []int{1, 2, 3},
			n:    5,
			want: true,
		},
	}

	for _, tc := range testcases {
		got := checkSubarraySum(tc.arr, tc.n)
		assert.Equal(t, tc.want, got)
	}
}
