package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSearchRange(t *testing.T) {
	myt := easytest.New(t).F(searchRange)
	myt.Input([]int{5, 7, 7, 8, 8, 10}, 8).Output([]int{3, 4})
	myt.Input([]int{5, 7, 7, 8, 8, 10}, 6).Output([]int{-1, -1})
	myt.Input([]int{}, 0).Output([]int{-1, -1})

	myt.Input([]int{1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7}, 4).Output([]int{3, 11})
}

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	left, right, mid := 0, len(nums)-1, 0

	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			// if mid == 0 || nums[mid-1] != target {
			// 	res[0] = mid
			// } else {
			// 	res[0] = searchRange(nums[left:mid], target)[0] + left
			// }
			// if mid == len(nums)-1 || nums[mid+1] != target {
			// 	res[1] = mid
			// } else {
			// 	res[1] = searchRange(nums[mid+1:right+1], target)[1] + mid + 1
			// }
			if mid > 0 && nums[mid-1] == target {
				res[0] = leftMost(nums[left:mid], target) + left
			} else {
				res[0] = mid
			}
			if mid < len(nums)-1 && nums[mid+1] == target {
				res[1] = rightMost(nums[mid+1:right+1], target) + mid + 1
			} else {
				res[1] = mid
			}
			return res
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func leftMost(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0

	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func rightMost(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0

	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
