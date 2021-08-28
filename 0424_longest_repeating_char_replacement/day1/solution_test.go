package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tcs := []struct {
		str  string
		k    int
		want int
	}{
		{
			str:  "AAABCAAAA",
			k:    1,
			want: 8,
		},
	}

	for _, tc := range tcs {
		assert.Equal(t, tc.want, characterReplacement(tc.str, tc.k))
	}
}
