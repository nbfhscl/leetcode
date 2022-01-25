package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIsSymmetric(t *testing.T) {
	myt := easytest.New(t).F(isSymmetric)
	myt.Input(NewTree(nil, 1, 2, 2)).Output(true)
	myt.Input(NewTree(nil, 1, 2, 2, 3, 4, 4, 3)).Output(true)
	myt.Input(NewTree(nil, 1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5)).Output(true)

	myt.Input(NewTree(nil, 1, 2, 3)).Output(false)
}

/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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

func isSymmetric(root *TreeNode) bool {
	queue := make([]*TreeNode, 0, 16)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i] == nil && queue[size-i-1] == nil {
				continue
			}
			if queue[i] == nil || queue[size-i-1] == nil {
				return false
			}
			if queue[i].Val != queue[size-i-1].Val {
				return false
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[size:]
	}
	return true
}

// @lc code=end

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isOutput(root.Left, root.Right)
}

func isOutput(tn1, tn2 *TreeNode) bool {
	if tn1 == nil && tn2 == nil {
		return true
	}
	if tn1 == nil || tn2 == nil {
		return false
	}

	return tn1.Val == tn2.Val && isOutput(tn1.Left, tn2.Right) && isOutput(tn1.Right, tn2.Left)
}
