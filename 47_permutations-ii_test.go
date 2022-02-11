package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestPermuteUnique(t *testing.T) {
	myt := easytest.New(t).F(permuteUnique)
	myt.Input([]int{1, 1, 2}).Output([][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}})
	myt.Input([]int{1, 2, 3}).Output([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}})
	// 1,1,2,2 -> 1,2,1,2 -> 1,2,2,1 -> 1,2,2,1 -> 1,2,1,2dup
	myt.Input([]int{2, 2, 1, 1}).Output([][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}})
	// myt.Input([]int{-1,2,0,-1,1,0,1}).Output([][]int)
}

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	backtracking_permuteUnique(&res, nums, []int{})
	return res
}

func backtracking_permuteUnique(res *[][]int, nums []int, ans []int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, ans...))
	} else {
		for i := range nums {
			if i == 0 || nums[i] != nums[i-1] {
				nums2 := append(append([]int{}, nums[:i]...), nums[i+1:]...)
				backtracking_permuteUnique(res, nums2, append(ans, nums[i]))
			}
		}
	}
}

// @lc code=end

func permuteUnique1(nums []int) [][]int {
	res := make([][]int, 0)
	// sort.Ints(nums)
	backtracking_permuteUnique1(&res, make([]int, 0), nums, 0)
	return res
}

func backtracking_permuteUnique1(res *[][]int, this, nums []int, k int) {
	if k == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, this)
		*res = append(*res, temp)
	} else {
		used := make(map[int]bool)
		for i := k; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			nums[k], nums[i] = nums[i], nums[k]
			this = append(this, nums[k])
			backtracking_permuteUnique1(res, this, nums, k+1)
			this = this[:len(this)-1]
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
}

func permuteUnique1_opt(nums []int) [][]int {
	res := make([][]int, 0)
	// sort.Ints(nums)
	backtracking_permuteUnique1_opt(&res, nums, 0)
	return res
}

func backtracking_permuteUnique1_opt(res *[][]int, nums []int, k int) {
	if k == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
	} else {
		used := make(map[int]struct{})
		for i := k; i < len(nums); i++ {
			if _, ok := used[nums[i]]; ok {
				continue
			}
			used[nums[i]] = struct{}{}
			nums[k], nums[i] = nums[i], nums[k]
			backtracking_permuteUnique1_opt(res, nums, k+1)
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
}

func permuteUnique2(nums []int) [][]int {
	return divideAndConquer(nums)
}

func divideAndConquer(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	} else {
		used := make(map[int]bool)
		remain := make([]int, len(nums)-1)
		res := make([][]int, 0)
		for i := range nums {
			if used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			copy(remain[:i], nums[:i])
			copy(remain[i:], nums[i+1:])
			sub := divideAndConquer(remain)
			for j := range sub {
				sub[j] = append(sub[j], nums[i])
			}
			res = append(res, sub...)
		}
		return res
	}
}
