package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  string
		k    int
		want string
	}{
		{
			str:  "1432219",
			k:    3,
			want: "1219",
		},
	}

	for _, tc := range testcases {
		got := removeKdigits(tc.str, tc.k)
		assert.Equal(t, tc.want, got)
	}
}
