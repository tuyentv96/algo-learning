package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2(t *testing.T) {

	tcs := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 3, 3, 3, 4},
			want: 9,
		},
		{
			nums: []int{10000},
			want: 10000,
		},
	}

	for _, tc := range tcs {
		got := deleteAndEarn(tc.nums)
		assert.Equalf(t, tc.want, got, "want: %d, got: %d", tc.want, got)
	}
}
