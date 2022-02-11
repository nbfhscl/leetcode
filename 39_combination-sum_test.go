package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCombinationSum(t *testing.T) {
	myt := easytest.New(t).F(combinationSum)
	myt.Input([]int{2, 3, 6, 7}, 7).Output([][]int{{2, 2, 3}, {7}})
}

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtracking_combinationSum(&res, candidates, []int{}, target, len(candidates)-1)
	return res
}

func backtracking_combinationSum(res *[][]int, nums, ans []int, target int, next int) {
	if target == 0 {
		*res = append(*res, append([]int{}, ans...))
	} else {
		for i := next; i >= 0; i-- {
			if target-nums[i] >= 0 {
				backtracking_combinationSum(res, nums, append(ans, nums[i]), target-nums[i], i)
			}
		}
	}
}

// @lc code=end
