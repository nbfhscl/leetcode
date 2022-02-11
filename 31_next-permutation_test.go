package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestNextPermutation(t *testing.T) {
	myt := easytest.New(t).F(nextPermutation)
	myt.Input([]int{1, 2, 3}).Output([]int{1, 3, 2})
	myt.Input([]int{1, 1, 5}).Output([]int{1, 5, 1})
	myt.Input([]int{3, 2, 1}).Output([]int{1, 2, 3})
	myt.Input([]int{1}).Output([]int{1})

	myt.Input([]int{1, 100, 89}).Output([]int{89, 1, 100})

	myt.Input([]int{1, 3, 2}).Output([]int{2, 1, 3})
	myt.Input([]int{2, 3, 1}).Output([]int{3, 1, 2})
	myt.Input([]int{8, 4, 6, 5, 4, 3, 2, 1}).Output([]int{8, 5, 1, 2, 3, 4, 4, 6})
}

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) []int {
	var k int
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			k = i
			break
		}
	}

	if k != 0 {
		for i := len(nums) - 1; i >= k; i-- {
			if nums[i] > nums[k-1] {
				nums[i], nums[k-1] = nums[k-1], nums[i]
				break
			}
		}
	}

	for i, j := len(nums)-1, k; i > j; i, j = i-1, j+1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// @lc code=end
