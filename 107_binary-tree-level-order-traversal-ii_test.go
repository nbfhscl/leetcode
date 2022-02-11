package test

import (
	"container/list"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestLevelOrderBottom(t *testing.T) {
	myt := easytest.New(t).F(levelOrderBottom)
	myt.Input(NewTree(nil, 3, 9, 20, nil, nil, 15, 7)).Output([][]int{{15, 7}, {9, 20}, {3}})
}

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	r := make([][]int, 0, 4)
	row := 0
	if root == nil {
		return r
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		r = append(r, make([]int, len(q)))
		for i, tn := range q {
			q = q[1:]
			r[row][i] = tn.Val.(int)
			if tn.Left != nil {
				q = append(q, tn.Left)
			}
			if tn.Right != nil {
				q = append(q, tn.Right)
			}
		}
		row++
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

// @lc code=end
func levelOrderBottom1(root *TreeNode) [][]int {
	r := make([][]int, 0, 4)
	row := 0
	q := list.New()
	if root == nil {
		return r
	}
	q.PushBack(root)
	for q.Len() > 0 {
		size := q.Len()
		r = append(r, make([]int, size))
		for i := 0; i < size; i++ {
			v := q.Front()
			q.Remove(v)
			tn := v.Value.(*TreeNode)
			r[row][i] = tn.Val.(int)
			if tn.Left != nil {
				q.PushBack(tn.Left)
			}
			if tn.Right != nil {
				q.PushBack(tn.Right)
			}
		}
		row++
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}
