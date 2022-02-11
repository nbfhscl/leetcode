package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestHasCycle(t *testing.T) {
	myt := easytest.New(t).F(hasCycle)
	myt.Run(NewListCycled(0, 0, 1)).Equal(true)
	myt.Run(NewListCycled(1, 0, 1, 2, 3, 4, 5, 6)).Equal(true)
	myt.Run(NewListCycled(-1, 0, 1, 2, 3, 4, 5, 6)).Equal(false)
	myt.Run(NewListCycled(0, 0)).Equal(true)
}

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	for i, j := head.Next, head.Next.Next; ; i, j = i.Next, j.Next.Next {
		if i == j {
			return true
		}
		if i.Next == nil || j.Next == nil || j.Next.Next == nil {
			break
		}
	}
	return false
}

// @lc code=end
func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return true
		} else {
			m[head] = true
		}
		head = head.Next
	}
	return false
}
