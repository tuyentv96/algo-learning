package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		connections [][]int
		n           int
		want        int
	}{
		{
			connections: [][]int{{1, 2, 5}, {1, 3, 6}, {2, 3, 1}},
			n:           3,
			want:        6,
		},
	}

	for _, tc := range testcases {
		got := minimumCost(tc.n, tc.connections)
		assert.Equal(t, tc.want, got)
	}
}
