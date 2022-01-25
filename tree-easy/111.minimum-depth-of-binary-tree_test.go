package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMinDepth(t *testing.T) {
	myt := easytest.New(t).F(minDepth)
	myt.Input(NewTree(nil, 3, 9, 20, nil, nil, 15, 7)).Output(2)
	myt.Input(NewTree(nil, 2, nil, 3, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, 5, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 6)).Output(5)
}

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var md int = 1
	q := []*TreeNode{root}
	for len(q) > 0 {
		for _, v := range q {
			q = q[1:]
			if v.Left == nil && v.Right == nil {
				return md
			}
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		md++
	}
	return md
}

// @lc code=end
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	ld := minDepth(root.Left)
	rd := minDepth(root.Right)
	if ld < rd {
		ld, rd = rd, ld
	}
	if rd == 0 {
		return ld + 1
	}
	return rd + 1
}
