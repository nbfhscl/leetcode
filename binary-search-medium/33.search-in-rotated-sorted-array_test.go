package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSearch(t *testing.T) {
	myt := easytest.New(t).F(search)
	myt.Input([]int{4, 5, 6, 7, 0, 1, 2}, 0).Output(4)
	myt.Input([]int{4, 5, 6, 7, 0, 1, 2}, 3).Output(-1)
	myt.Input([]int{1}, 0).Output(-1)

	myt.Input([]int{1, 3}, 0).Output(-1)
	myt.Input([]int{1, 3}, 4).Output(-1)
	myt.Input([]int{1, 3}, 1).Output(0)
	myt.Input([]int{1, 3}, 3).Output(1)
	myt.Input([]int{3, 1}, 0).Output(-1)
	myt.Input([]int{3, 1}, 4).Output(-1)
	myt.Input([]int{3, 1}, 1).Output(1)
	myt.Input([]int{3, 1}, 3).Output(0)
}

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end
