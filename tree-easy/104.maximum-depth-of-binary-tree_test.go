package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMaxDepth(t *testing.T) {
	myt := easytest.New(t).F(maxDepth)
	myt.Input(NewTree(nil, 3, 9, 20, nil, nil, 15, 7)).Output(3)
	myt.Input(NewTree(nil, 1, nil, 2)).Output(2)
	myt.Input(NewTree(nil)).Output(0)
	myt.Input(NewTree(nil, 0)).Output(1)
}

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	if ld > rd {
		return ld + 1
	}

	return rd + 1
}

// @lc code=end

func maxDepth1(root *TreeNode) int {
	var max int
	depth(root, &max, 0)
	return max
}

func depth(tn *TreeNode, max *int, v int) {
	if tn == nil {
		return
	}
	v++
	if v > *max {
		*max = v
	}
	depth(tn.Left, max, v)
	depth(tn.Right, max, v)
}

func maxDepth2(root *TreeNode) int {
	q := make([]*TreeNode, 0, 4)
	q = append(q, root)
	depth := 0
	for len(q) > 0 {
		ok := false
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i] != nil {
				ok = true
				q = append(q, q[i].Left, q[i].Right)
			}
		}
		if !ok {
			return depth
		}
		depth++
		q = q[size:]
	}
	return depth
}
