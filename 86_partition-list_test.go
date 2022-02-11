package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestPartition(t *testing.T) {
	myt := easytest.New(t).F(partition)
	myt.Input(NewList(1, 4, 3, 2, 5, 2), 3).Output(NewList(1, 2, 2, 4, 3, 5))
	myt.Input(NewList(2, 1), 2).Output(NewList(1, 2))
	myt.Input(NewList(1, 4, 3, 0, 2, 5, 2), 3).Output(NewList(1, 0, 2, 2, 4, 3, 5))
}

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hh := &ListNode{}
	hh.Next = head
	p1 := hh
	for p1.Next != nil && p1.Next.Val < x {
		p1 = p1.Next
	}
	if p1.Next == nil {
		return head
	}
	p2 := p1.Next
	for p2.Next != nil {
		if p2.Next.Val < x {
			p1.Next, p2.Next.Next, p2.Next = p2.Next, p1.Next, p2.Next.Next
			p1 = p1.Next
		} else {
			p2 = p2.Next
		}
	}
	return hh.Next
}

// @lc code=end
