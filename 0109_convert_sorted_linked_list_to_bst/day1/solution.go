package day1

import (
	"github.com/tuyentv96/algo-learning/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	return buildBST(head, nil)
}

func buildBST(head *ListNode, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	slow := head
	fast := head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var root TreeNode
	root.Val = slow.Val
	root.Left = buildBST(head, slow)
	root.Right = buildBST(slow.Next, tail)
	return &root
}

//func sortedListToBST(head *ListNode) *TreeNode {
//	if head == nil {
//		return nil
//	}
//	if head != nil && head.Next == nil {
//		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
//	}
//	middleNode, preNode := middleNodeAndPreNode(head)
//	if middleNode == nil {
//		return nil
//	}
//	if preNode != nil {
//		preNode.Next = nil
//	}
//	if middleNode == head {
//		head = nil
//	}
//	return &TreeNode{Val: middleNode.Val, Left: sortedListToBST(head), Right: sortedListToBST(middleNode.Next)}
//}
//
//func middleNodeAndPreNode(head *ListNode) (middle *ListNode, pre *ListNode) {
//	if head == nil || head.Next == nil {
//		return nil, head
//	}
//	p1 := head
//	p2 := head
//	for p2.Next != nil && p2.Next.Next != nil {
//		pre = p1
//		p1 = p1.Next
//		p2 = p2.Next.Next
//	}
//	return p1, pre
//}

//func sortedListToBST(head *ListNode) *TreeNode {
//	return buildBST(head)
//}
//
//func buildBST(head *ListNode) *TreeNode {
//	mid := getMid(head)
//	if mid == nil {
//		return nil
//	}
//
//	if mid == head {
//		return nil
//	}
//
//	var root TreeNode
//	root.Val = mid.Val
//	root.Left = (*structures.TreeNode)(buildBST(head))
//	root.Right = (*structures.TreeNode)(buildBST(mid))
//	return &root
//}
//
//func getMid(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		fast.Next = fast.Next.Next
//		slow = (*ListNode)(slow.Next)
//	}
//
//	return slow
//}
