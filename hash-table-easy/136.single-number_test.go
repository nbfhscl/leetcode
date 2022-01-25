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
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] {
			delete(m, v)
		} else {
			m[v] = true
		}
	}
	for k := range m {
		return k
	}
	return -1
}

// @lc code=end
