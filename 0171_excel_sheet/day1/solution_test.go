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
			str:  "A",
			want: 1,
		},
		{
			str:  "AA",
			want: 27,
		},
	}

	for _, tc := range testcases {
		got := titleToNumber(tc.str)
		assert.Equal(t, tc.want, got)
	}
}
