package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIsSameTree(t *testing.T) {
	myt := easytest.New(t).F(isSameTree)
	myt.Input(NewTree(nil, 1, 2, 3), NewTree([]interface{}{nil, 1, 2, 3})).Output(true)
	myt.Input(NewTree(nil, 1, 2, 3, 4, 5, 6, 7), NewTree([]interface{}{nil, 1, 2, 3, 4, 5, 6, 7})).Output(true)
	myt.Input(NewTree(nil, 1, 2, 4, 5, 6, 7), NewTree([]interface{}{nil, 1, 2, 3, 4, 5, 6, 7})).Output(false)

}

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if ok := isSameTree(p.Left, q.Left); !ok {
		return false
	}
	if ok := isSameTree(p.Right, q.Right); !ok {
		return false
	}
	return true
}

// @lc code=end
