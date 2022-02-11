package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestPermute(t *testing.T) {
	myt := easytest.New(t).F(permute)
	// 1,2,3 -> 1,3,2 -> 2,1,3 -> 2,3,1 -> 3,2,1 -> 3,1,2
	myt.Input([]int{1, 2, 3}).Output([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}})
}

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	// sort.Ints(nums)
	backtracking_permute(&res, nums, 0)
	return res
}

func backtracking_permute(res *[][]int, nums []int, k int) {
	if k == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	} else {
		for i := k; i < len(nums); i++ {
			nums[k], nums[i] = nums[i], nums[k]
			backtracking_permute(res, nums, k+1)
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
}

// @lc code=end
