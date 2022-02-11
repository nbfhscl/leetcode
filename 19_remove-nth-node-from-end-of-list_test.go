package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestRemoveNthFromEnd(t *testing.T) {
	myt := easytest.New(t).F(removeNthFromEnd)
	myt.Input(NewList(1, 2, 3, 4, 5), 2).Output(NewList(1, 2, 3, 5))
	myt.Input(NewList(1), 1).Output(NewList())
	myt.Input(NewList(1, 2), 1).Output(NewList(1))
	myt.Input(NewList(1, 2, 3), 3).Output(NewList(2, 3))
	myt.Input(NewList(1, 2, 3), 2).Output(NewList(1, 3))
	myt.Input(NewList(1, 2, 3), 1).Output(NewList(1, 2))
}

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var i, j *ListNode = head, head
	for n > 0 && i.Next != nil {
		i = i.Next
		n--
	}
	if n == 1 {
		return head.Next
	}
	for i.Next != nil {
		i = i.Next
		j = j.Next
	}
	j.Next = j.Next.Next
	return head
}

// @lc code=end
