package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSortedArrayToBST(t *testing.T) {
	myt := easytest.New(t).F(sortedArrayToBST)
	myt.Input([]int{-10, -3, 0, 5, 9}).Output(NewTree(nil, 0, -3, 9, -10, nil, 5))
}

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	tn := &TreeNode{Val: nums[mid]}
	if mid != 0 {
		tn.Left = sortedArrayToBST(nums[:mid])
		tn.Right = sortedArrayToBST(nums[mid+1:])
	}
	return tn
}

// @lc code=end
