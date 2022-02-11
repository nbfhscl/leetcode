package test

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	node := new(ListNode)
	node.Val = 9
	l1.Next = node
	node = new(ListNode)
	node.Val = 11
	node.Next = nil
	l1.Next.Next = node

	l2 := new(ListNode)
	l2.Val = 4
	node = new(ListNode)
	node.Val = 33
	l2.Next = node
	node = new(ListNode)
	node.Val = 44
	node.Next = nil
	l2.Next.Next = node

	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		t.Logf("%v - ", l3.Val)
		l3 = l3.Next
	}
	t.Logf("nil")
}

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for {
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next
	}
	return head.Next
}

// @lc code=end
