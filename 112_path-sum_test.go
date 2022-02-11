package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestHasPathSum(t *testing.T) {
	myt := easytest.New(t).F(hasPathSum)
	myt.Input(NewTree(nil, 5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, nil, nil, nil, 1), 22).Output(true)
	myt.Input(NewTree(nil, 1, 2, 3), 5).Output(false)
	myt.Input(NewTree(nil, 1, 2), 5).Output(false)
	myt.Input(NewTree(nil), 0).Output(false)
	myt.Input(NewTree(nil, 1, 2, 3, 2, 1, 1, 1), 5).Output(true)

	// must be root-to-leaf path not only root
	//     1
	//    2
	myt.Input(NewTree(nil, 1, 2), 1).Output(false)
	//     1
	//   -2 3
	myt.Input(NewTree(nil, 1, -2, 3), 1).Output(false)
	// if there is only one node, it is a root and also a leaf
	myt.Input(NewTree(nil, 1), 1).Output(true)
	// 1,2,nil,3,nil,4,nil,5	6
	//		1
	//     2
	//    3
	//   4
	//  5
	myt.Input(NewTree(nil, 1, 2, nil, 3, nil, 4, nil, 5), 6).Output(false)

}

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// return true in and only in this situation
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	targetSum -= root.Val.(int)
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// @lc code=end
