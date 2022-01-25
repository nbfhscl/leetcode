package test

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(ns ...int) *ListNode {
	if len(ns) == 0 {
		return nil
	}
	ln := &ListNode{Val: ns[0]}
	head := ln
	for _, v := range ns[1:] {
		ln.Next = &ListNode{}
		ln = ln.Next
		ln.Val = v
	}
	ln.Next = nil
	return head
}

func NewList2(ns ...int) *ListNode {
	if len(ns) == 0 {
		return nil
	}
	ln := &ListNode{Val: ns[0]}
	if len(ns) > 1 {
		ln.Next = NewList2(ns[1:]...)
	} else {
		ln.Next = nil
	}
	return ln
}

func NewListCycled(pos int, ns ...int) *ListNode {
	if len(ns) == 0 {
		return nil
	}
	if pos < 0 {
		return NewList(ns...)
	}
	var lnpos *ListNode
	head := &ListNode{Val: ns[0]}
	if pos == 0 {
		lnpos = head
	}
	ln := head
	for i, v := range ns[1:] {
		ln.Next = &ListNode{Val: v}
		ln = ln.Next
		if i+1 == pos {
			lnpos = ln
		}
	}
	ln.Next = lnpos
	return head
}
func NewListIntersected(ns1 []int, ns2 []int, ns3 []int) (*ListNode, *ListNode) {
	a, b, c := NewList(ns1...), NewList(ns2...), NewList(ns3...)
	aa, bb := a, b
	if aa != nil {
		for aa.Next != nil {
			aa = aa.Next
		}
		aa.Next = c
	} else {
		a = c
	}
	if bb != nil {
		for bb.Next != nil {
			bb = bb.Next
		}
		bb.Next = c
	} else {
		b = c
	}
	return a, b
}
