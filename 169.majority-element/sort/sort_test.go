package sort_test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestQuickSort(t *testing.T) {
	myt := easytest.New(t).F(QuickSort_Slice)
	myt.Run([]int{1, 2, 3, 4, 5, 6, 7, 8}).Equal([]int{1, 2, 3, 4, 5, 6, 7, 8})
}

func TestMajorityElement(t *testing.T) {
	myt := easytest.New(t).F(majorityElement)
	myt.Run([]int{3, 2, 3}).Equal(3)
	myt.Run([]int{2, 2, 1, 1, 1, 2, 2}).Equal(2)
	myt.Run([]int{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 1}).Equal(1)
}

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start

// sort
func majorityElement(nums []int) int {
	QuickSort_Slice(nums)
	return nums[len(nums)/2]
}

func QuickSort_Slice(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	doPivot_Slice(nums)
	p := partition(nums)
	QuickSort_Slice(nums[:p])
	QuickSort_Slice(nums[p+1:])
	return nums
}

func doPivot_Slice(nums []int) {
	mid := len(nums) / 2
	if nums[mid] < nums[0] {
		nums[0], nums[mid] = nums[mid], nums[0]
	}
	if nums[mid] < nums[len(nums)-1] {
		nums[mid], nums[len(nums)-1] = nums[len(nums)-1], nums[mid]
	}
	if nums[0] < nums[len(nums)-1] {
		nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
	}
}

func partition(nums []int) int {
	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			p++
			nums[i], nums[p] = nums[p], nums[i]
		}
	}
	nums[0], nums[p] = nums[p], nums[0]
	return p
}

// @lc code=end
