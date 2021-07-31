package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		intervals [][]int
		new       []int
		want      [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {6, 9}},
			new:       []int{2, 5},
			want:      [][]int{{1, 5}, {6, 9}},
		},
	}

	for _, tc := range testcases {
		got := insert(tc.intervals, tc.new)
		assert.Equal(t, tc.want, got)
	}
}
