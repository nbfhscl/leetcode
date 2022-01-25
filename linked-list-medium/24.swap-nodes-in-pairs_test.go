package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSwapPairs(t *testing.T) {
	myt := easytest.New(t).F(swapPairs)
	myt.Input(NewList(1, 2, 3, 4)).Output(NewList(2, 1, 4, 3))
	myt.Input(NewList()).Output(NewList())
	myt.Input(NewList(1)).Output(NewList(1))
	myt.Input(NewList(1, 2, 3)).Output(NewList(2, 1, 3))
	myt.Input(NewList(1, 2)).Output(NewList(2, 1))
}

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	for i := h; i.Next != nil && i.Next.Next != nil; i.Next.Next.Next, i.Next.Next, i.Next, i = i.Next, i.Next.Next.Next, i.Next.Next, i.Next {
	}
	return h.Next
}

// @lc code=end
