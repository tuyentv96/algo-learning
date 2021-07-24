package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{
			str:  "pwwkew",
			want: 3,
		},
		{
			str:  "abcabcbb",
			want: 3,
		},
	}

	for _, tc := range testcases {
		got := lengthOfLongestSubstring(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
