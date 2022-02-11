package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIsBalanced(t *testing.T) {
	myt := easytest.New(t).F(isBalanced)
	myt.Input(NewTree(nil, 1, 2, 2, 3, 3, nil, nil, 4, 4)).Output(false)
	myt.Input(NewTree(nil, 3, 9, 20, nil, nil, 15, 7)).Output(true)
}

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	md := maxDepthForIsBalanced(root)
	if md == -1 {
		return false
	}
	return true
}

func maxDepthForIsBalanced(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	if tn.Left == nil && tn.Right == nil {
		return 1
	}
	ld := maxDepthForIsBalanced(tn.Left)
	if ld == -1 {
		return -1
	}
	rd := maxDepthForIsBalanced(tn.Right)
	if rd == -1 {
		return -1
	}
	if ld < rd {
		ld, rd = rd, ld
	}
	if ld-rd > 1 {
		return -1
	} else {
		return ld + 1
	}
}

// @lc code=end
