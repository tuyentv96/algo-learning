package day1

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	testcases := []struct {
		l1   *ListNode
		l2   *ListNode
		want []int
	}{
		{
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			want: []int{7, 0, 8},
		},
	}

	for _, tc := range testcases {
		got := addTwoNumbers(tc.l1, tc.l2)
		assert.NoError(t, check(got, tc.want))
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
