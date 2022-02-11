package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSingleNumber(t *testing.T) {
	myt := easytest.New(t).F(singleNumber)
	myt.Run([]int{2, 2, 1}).Equal(1)
	myt.Run([]int{4, 1, 2, 1, 2}).Equal(4)
	myt.Run([]int{1}).Equal(1)
}

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	var r int
	for _, v := range nums {
		r ^= v
	}
	return r
}

// @lc code=end
