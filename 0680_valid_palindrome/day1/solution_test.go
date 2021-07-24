package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  string
		want bool
	}{
		{
			str:  "aba",
			want: true,
		},
	}

	for _, tc := range testcases {
		got := validPalindrome(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
