package day1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testcases := []struct {
		arr  []int
		want []int
	}{
		{
			arr:  []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
		{
			arr:  []int{7, 8, 3, 6, 4, 2, 1},
			want: []int{7, 8, 4, 1, 2, 3, 6},
		},
		{
			arr:  []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			arr:  []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
	}

	for _, tc := range testcases {
		time.Sleep(time.Hour)
		nextPermutation(tc.arr)
		assert.Equal(t, tc.want, tc.arr)
	}
}
