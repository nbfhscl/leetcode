package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSortColor(t *testing.T) {
	myt := easytest.New(t).F(sortColors)
	myt.Input([]int{2, 0, 2, 1, 1, 0}).Output([]int{0, 0, 1, 1, 2, 2})
	myt.Input([]int{2, 0, 1}).Output([]int{0, 1, 2})
	myt.Input([]int{0}).Output([]int{0})
	myt.Input([]int{1}).Output([]int{1})
}

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		k := nums[i]
		j := i
		for ; j > 0 && nums[j-1] > k; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = k
	}
	return nums
}

// @lc code=end
