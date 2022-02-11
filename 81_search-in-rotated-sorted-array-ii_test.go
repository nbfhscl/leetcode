package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSearchRotatedSortedArray2(t *testing.T) {
	myt := easytest.New(t).F(searchRotatedSortedArray2)
	myt.Input([]int{2, 5, 6, 0, 0, 1, 2}, 0).Output(true)
	myt.Input([]int{2, 5, 6, 0, 0, 1, 2}, 3).Output(false)
	myt.Input([]int{2, 2, 2, 2, 2}, 2).Output(true)
	myt.Input([]int{2, 2, 2, 2, 2, 3}, 3).Output(true)
	myt.Input([]int{3, 2, 2, 2, 2, 2}, 3).Output(true)
	myt.Input([]int{1, 0, 1, 1, 1}, 0).Output(true)
	// 应该取左边区间
	myt.Input([]int{1, 1, 2, 3, 4, 1, 1, 1, 1, 1, 1, 1, 1}, 4).Output(true)
	// 应该取右边区间
	myt.Input([]int{1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 1, 1}, 4).Output(true)
	// 取哪边都可以
	myt.Input([]int{1, 1, 1}, 4).Output(false)
	myt.Input([]int{9, 2, 9}, 2).Output(true)
	myt.Input([]int{9, 2, 9}, 9).Output(true)
	myt.Input([]int{9, 2, 9}, 8).Output(false)
}

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func searchRotatedSortedArray2(nums []int, target int) bool {
	// func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left++
		}
	}
	return false
}

// @lc code=end
