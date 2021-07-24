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
			str:  ".,",
			want: true,
		},
		{
			str:  "A man, a plan, a canal: Panama",
			want: true,
		},
	}

	for _, tc := range testcases {
		got := isPalindrome(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
