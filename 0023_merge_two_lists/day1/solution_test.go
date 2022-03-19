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
				Val: -9,
				Next: &ListNode{
					Val: 3,
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
				},
			},
			want: []int{-9, 3, 5, 7},
		},
		{
			l1:   nil,
			l2:   nil,
			want: []int{},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}

	for _, tc := range testcases {
		got := mergeTwoLists(tc.l1, tc.l2)
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
