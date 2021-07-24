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
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
	}

	for _, tc := range testcases {
		got := maxProfit(tc.prices)
		assert.Equal(t, tc.want, got)
	}
}
