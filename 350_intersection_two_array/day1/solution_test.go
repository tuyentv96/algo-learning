package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		//{
		//	nums1: []int{1, 2, 2, 1},
		//	nums2: []int{2, 2},
		//	want:  []int{2, 2},
		//},
		{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
	}

	for _, tc := range testcases {
		got := intersect(tc.nums1, tc.nums2)
		assert.Equal(t, tc.want, got)
	}
}
