package day2

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testcases := []struct {
		l1   *ListNode
		want []int
	}{
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			want: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tc := range testcases {
		reorderList(tc.l1)
		assert.NoError(t, check(tc.l1, tc.want))
	}
}

func check(l *ListNode, arr []int) error {
	i := 0
	for l != nil {
		if l.Val != arr[i] {
			return errors.New(fmt.Sprintf("want: %d, got: %d", arr[i], l.Val))
		}

		l = l.Next
		i++
	}

	return nil
}
