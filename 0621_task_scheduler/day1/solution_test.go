package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		str  []byte
		n    int
		want int
	}{
		{
			str:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:    2,
			want: 8,
		},
	}

	for _, tc := range testcases {
		got := leastInterval(tc.str, tc.n)
		assert.Equal(t, tc.want, got)
	}
}
