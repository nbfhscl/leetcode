package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGetIntersectionNode(t *testing.T) {
	myt := easytest.New(t).F(getIntersectionNode)
	myt.Equal(getIntersectionNode(NewListIntersected([]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5})).Val, 8)
	myt.Equal(getIntersectionNode(NewListIntersected([]int{1, 9, 1}, []int{3}, []int{2, 4})).Val, 2)
	myt.Equal(getIntersectionNode(NewListIntersected([]int{2, 6, 4}, []int{1, 5}, []int{})), nil)
	// myt.Equal(getIntersectionNode(NewListIntersected([]int{}, []int{}, []int{2, 4})), nil)
	myt.Equal(getIntersectionNode(NewListIntersected([]int{1, 2, 3}, []int{}, []int{4, 5})), nil)
}

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	if pa == nil || pb == nil {
		return nil
	}
	for pa != pb {
		pa, pb = pa.Next, pb.Next
		if pa == pb {
			return pa
		}
		if pa == nil {
			pa = headB
		}
		if pb == nil {
			pb = headA
		}
	}
	// if pa == pb, return pa rather than nil
	return pa
}

// @lc code=end

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	var la, lb int
	a, b := headA, headB
	for a != nil {
		la++
		a = a.Next
	}
	for b != nil {
		lb++
		b = b.Next
	}
	a, b = headA, headB
	for la > lb {
		a = a.Next
		la--
	}
	for lb > la {
		b = b.Next
		lb--
	}
	for a != nil {
		if a == b {
			return a
		}
		a, b = a.Next, b.Next
	}
	return nil

}
