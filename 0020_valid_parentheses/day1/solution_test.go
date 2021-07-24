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
			str:  "{}",
			want: true,
		},
		{
			str:  "()[]{}",
			want: true,
		},
		{
			str:  "(]",
			want: false,
		},
	}

	for _, tc := range testcases {
		got := isValid(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
