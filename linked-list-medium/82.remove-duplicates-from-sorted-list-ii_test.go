package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestDeleteDuplicates(t *testing.T) {
	myt := easytest.New(t).F(deleteDuplicates)
	myt.Input(NewList(1, 1, 1, 2, 3)).Output(NewList(2, 3))
	myt.Input(NewList(1, 2, 3, 3, 4, 4, 5)).Output(NewList(1, 2, 5))
}

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hh := &ListNode{}
	hh.Next = head

	p1 := hh
	for p1.Next != nil {
		p2 := p1.Next
		var dup bool
		for p2.Next != nil && p2.Next.Val == p2.Val {
			p2 = p2.Next
			dup = true
		}
		if dup {
			p1.Next = p2.Next
		} else {
			p1 = p2
		}
	}
	return hh.Next
}

// @lc code=end

// 1 1 1 2 3
// h   p
// 2 3
// h

// 1 2 3
// hp
// 1 2 3
// h

// 1 1 1
// h   p
// nil
// h
