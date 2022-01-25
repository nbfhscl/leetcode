package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestAddTwoNumbers(t *testing.T) {
	myt := easytest.New(t).F(addTwoNumbers)
	myt.Input(NewList(2, 4, 3), NewList(5, 6, 4)).Output(NewList(7, 0, 8))
	myt.Input(NewList(0), NewList(5, 6, 4)).Output(NewList(5, 6, 4))
	myt.Input(NewList(0), NewList(0)).Output(NewList(0))
	myt.Input(NewList(9, 9), NewList(1, 1)).Output(NewList(0, 1, 1))
	myt.Input(NewList(9, 9, 9, 9, 9, 9, 9), NewList(9, 9, 9, 9)).Output(NewList(8, 9, 9, 9, 0, 0, 0, 1))
}

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := new(ListNode)
	for ln, c := r, 0; l1 != nil || l2 != nil || c != 0; ln = ln.Next {
		if l1 != nil {
			c += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			c += l2.Val
			l2 = l2.Next
		}
		ln.Next = &ListNode{Val: c % 10}
		c /= 10
	}
	return r.Next
}

// @lc code=end
