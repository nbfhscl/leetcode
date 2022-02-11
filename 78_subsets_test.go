package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSubsets(t *testing.T) {
	myt := easytest.New(t).F(subsets)
	myt.Input([]int{1, 2, 3}).Output([][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}})
}

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtracking_subsets(&res, nums, []int{}, 0)
	return res
}

func backtracking_subsets(res *[][]int, nums, ans []int, k int) {
	*res = append(*res, append([]int{}, ans...))
	for i := k; i < len(nums); i++ {
		backtracking_subsets(res, nums, append(ans, nums[i]), i+1)
	}
}

// @lc code=end
