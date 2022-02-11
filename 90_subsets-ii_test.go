package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSubsetsWithDup(t *testing.T) {
	myt := easytest.New(t).F(subsetsWithDup)
	myt.Input([]int{1, 2, 2}).Output([][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}})
}

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	backtracking_subsetsWithDup(&res, nums, make([]int, 0), 0)
	return res
}

func backtracking_subsetsWithDup(res *[][]int, nums, ans []int, k int) {
	*res = append(*res, append([]int{}, ans...))
	for i := k; i < len(nums); i++ {
		if i > k && nums[i] == nums[i-1] {
			continue
		}
		ans = append(ans, nums[i])
		backtracking_subsetsWithDup(res, nums, ans, i+1)
		ans = ans[:len(ans)-1]
	}
}

// @lc code=end
