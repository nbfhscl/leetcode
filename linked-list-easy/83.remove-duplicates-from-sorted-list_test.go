package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestDeleteDuplicates(t *testing.T) {
	myt := easytest.New(t)
	myt.F(deleteDuplicates).Run(NewList(1, 1, 2, 3, 3)).Equal(NewList(1, 2, 3))
	myt.Run(NewList(1, 1, 2)).Equal(NewList(1, 2))
	myt.Run(NewList(1)).Equal(NewList(1))
	myt.Run(NewList(-1, 0, 0, 100)).Equal(NewList(-1, 0, 100))
	myt.Run(NewList()).Equal(NewList())
}

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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
	if head == nil {
		return nil
	}
	r := head
	for r.Next != nil {
		if r.Val == r.Next.Val {
			r.Next = r.Next.Next
			continue
		}
		r = r.Next
	}
	return head
}

// @lc code=end
