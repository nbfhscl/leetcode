package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMajorityElement(t *testing.T) {
	myt := easytest.New(t).F(majorityElement)
	myt.Run([]int{3, 2, 3}).Equal(3)
	myt.Run([]int{3, 3, 4}).Equal(3)
	myt.Run([]int{3}).Equal(3)
	myt.Run([]int{2, 2, 1, 1, 1, 2, 2}).Equal(2)
	myt.Run([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 1}).Equal(1)
}

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start

func majorityElement(nums []int) int {
	return -1
}

// @lc code=end
