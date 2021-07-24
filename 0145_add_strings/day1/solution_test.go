package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str1 string
		str2 string
		want string
	}{
		{
			str1: "11",
			str2: "123",
			want: "134",
		},
	}

	for _, tc := range testcases {
		got := addStrings(tc.str1, tc.str2)
		assert.Equal(t, tc.want, got)
	}
}
