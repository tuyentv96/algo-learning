package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr    []int
		target int
		want   []int
	}{
		{
			arr:    []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want: []int{
				3, 4,
			},
		},
	}

	for _, tc := range testcases {
		got := searchRange(tc.arr, tc.target)
		assert.Equal(t, tc.want, got)
	}
}
