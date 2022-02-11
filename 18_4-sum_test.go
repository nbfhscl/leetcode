package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestFourSum(t *testing.T) {
	myt := easytest.New(t).F(fourSum)
	myt.Input([]int{1, 0, -1, 0, -2, 2}, 0).Output([][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}})
	myt.Input([]int{}, 0).Output([][]int{})
	myt.Input([]int{0, 0, 0, 0}, 0).Output([][]int{{0, 0, 0, 0}})
	myt.Input([]int{-2, -1, -1, 1, 1, 2, 2}, 0).Output([][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}})
}

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] == target-nums[i]-nums[j] {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] > target-nums[i]-nums[j] {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}

// @lc code=end
