package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{1, 2, 3, 1},
			want:   4,
		},
	}

	for _, tc := range testcases {
		got := rob(tc.prices)
		assert.Equal(t, tc.want, got)
	}
}
