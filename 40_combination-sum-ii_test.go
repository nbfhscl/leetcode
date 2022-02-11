package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCombinationSum2(t *testing.T) {
	myt := easytest.New(t).F(combinationSum2)
	myt.Input([]int{10, 1, 2, 6, 7, 1, 5}, 8).Output([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})
}

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtracking_combinationSum2(&res, candidates, []int{}, target, len(candidates)-1)
	return res
}

func backtracking_combinationSum2(res *[][]int, nums, ans []int, target int, next int) {
	if target == 0 {
		*res = append(*res, append([]int{}, ans...))
	} else {
		for i := next; i >= 0; i-- {
			if (i == next || nums[i] != nums[i+1]) && target-nums[i] >= 0 {
				backtracking_combinationSum2(res, nums, append(ans, nums[i]), target-nums[i], i-1)
			}
		}
	}
}

// @lc code=end
