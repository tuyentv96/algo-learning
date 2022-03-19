package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCS(t *testing.T) {
	tcs := []struct {
		text1 string
		text2 string
		want  int
	}{
		{
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			text1: "ylqpejqbalahwr",
			text2: "yrkzavgdmdgtqpg",
			want:  3,
		},
	}

	for _, tc := range tcs {
		got := longestCommonSubsequence(tc.text1, tc.text2)
		assert.Equal(t, tc.want, got, "want: %d, got: %d", tc.want, got)
	}
}
