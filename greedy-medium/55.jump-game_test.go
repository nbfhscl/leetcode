package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCanJump(t *testing.T) {
	myt := easytest.New(t).F(canJump)
	myt.Input([]int{2, 3, 1, 1, 4}).Output(true)
	myt.Input([]int{3, 2, 1, 0, 4}).Output(false)
	myt.Input([]int{1}).Output(true)
	myt.Input([]int{0}).Output(true)
	myt.Input([]int{0, 1}).Output(false)
	myt.Input([]int{1, 1, 1, 0}).Output(true)
}

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	now, max := 0, nums[0]
	for {
		more := false
		for i := now + 1; i <= max && i < len(nums); i++ {
			if nums[i]+i > max {
				max = nums[i] + i
				more = true
				now = i
			}
		}
		if !more {
			break
		}
	}
	if max >= len(nums)-1 {
		return true
	}
	return false
}

// @lc code=end
