package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{
			str:  "pwwkew",
			want: 4,
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

func Test2(t *testing.T) {
	testcases := []struct {
		str  string
		want int
	}{
		{
			str:  "pwwkew",
			want: 4,
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
