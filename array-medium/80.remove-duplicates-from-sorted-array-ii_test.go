package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestRemoveDuplicates(t *testing.T) {
	myt := easytest.New(t).F(removeDuplicates)
	myt.Input([]int{1, 1, 1, 2, 2, 3}).OutputWithInput([]int{1, 1, 2, 2, 3}, 5)
	myt.Input([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}).OutputWithInput([]int{0, 0, 1, 1, 2, 3, 3}, 7)
	myt.Input([]int{1, 2, 3, 3, 3}).OutputWithInput([]int{1, 2, 3, 3}, 4)
}

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var k int
	for k < len(nums)-1 && nums[k] == nums[k+1] {
		k++
	}
	if k > 0 {
		k--
	}
	for i := k + 2; i < len(nums); i++ {
		if nums[i] == nums[i-1] && nums[i] == nums[i-2] {
			for j := i - 2; j >= k+1; j-- {
				nums[j] = nums[j-1]
			}
			k++
		}
	}
	res := len(nums) - k
	copy(nums, nums[k:])
	return res
}

// @lc code=end
