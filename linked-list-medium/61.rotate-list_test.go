package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestRotateRight(t *testing.T) {
	myt := easytest.New(t).F(rotateRight)
	myt.Input(NewList(1, 2, 3, 4, 5), 0).Output(NewList(1, 2, 3, 4, 5))
	myt.Input(NewList(1, 2, 3, 4, 5), 1).Output(NewList(5, 1, 2, 3, 4))
	myt.Input(NewList(1, 2, 3, 4, 5), 2).Output(NewList(4, 5, 1, 2, 3))
	myt.Input(NewList(1, 2, 3, 4, 5), 3).Output(NewList(3, 4, 5, 1, 2))
}

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var length int = 1
	var p *ListNode = head
	for p.Next != nil {
		p = p.Next
		length++
	}
	var tail *ListNode = p
	k %= length
	if k == 0 {
		return head
	}

	p = head
	for length-1-k > 0 {
		p = p.Next
		k++
	}

	tail.Next = head
	head = p.Next
	p.Next = nil

	return head
}

// @lc code=end
