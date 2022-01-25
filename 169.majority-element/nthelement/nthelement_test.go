package nthelement_test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestNthElement(t *testing.T) {
	myt := easytest.New(t).F(nthElement)
	myt.Run([]int{1, 2, 3, 4, 5, 6, 7, 8}, 0, 7, 5).Equal(5)
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
	return nums[nthElement(nums, 0, len(nums)-1, len(nums)/2)]
}

func nthElement(nums []int, left, right int, n int) int {
	if left >= right {
		return left
	}
	doPivot(nums, left, right)
	p := partition(nums, left, right)
	if p == n {
		return p
	} else if p > n {
		return nthElement(nums, left, p-1, n)
	} else if p < n {
		return nthElement(nums, p+1, right, n)
	}
	return -1
}

func doPivot(nums []int, left, right int) {
	mid := left + (right-left)>>1
	if nums[mid] < nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	if nums[mid] < nums[left] {
		nums[mid], nums[left] = nums[left], nums[mid]
	}
	if nums[left] < nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func partition(nums []int, left, right int) int {
	p := left
	for i := left + 1; i <= right; i++ {
		if nums[i] < nums[left] {
			p++
			nums[i], nums[p] = nums[p], nums[i]
		}
	}
	nums[left], nums[p] = nums[p], nums[left]
	return p
}

// @lc code=end
