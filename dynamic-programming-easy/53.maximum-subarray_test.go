package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMaxSubArray(t *testing.T) {
	myt := easytest.New(t)
	myt.F(maxSubArray).Run([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}).Equal(6)
	myt.Run([]int{-2, -1, -3, -9}).Equal(-1)
	myt.Run([]int{-2, 1}).Equal(1)
	myt.Run([]int{-1}).Equal(-1)
}

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	prev, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if prev <= 0 {
			prev = nums[i]
		} else {
			prev += nums[i]
		}
		if max < prev {
			max = prev
		}
	}
	return max
}

// @lc code=end

func maxSubArray2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	var max int = dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
