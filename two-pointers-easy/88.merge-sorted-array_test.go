package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMerge(t *testing.T) {
	myt := easytest.New(t).F(merge)
	myt.Run([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3).Equal([]int{1, 2, 2, 3, 5, 6})
	myt.Run([]int{1}, 1, []int{}, 0).Equal([]int{1})
	myt.Run([]int{0}, 0, []int{1}, 1).Equal([]int{1})
}

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := m-1, n-1, m+n-1
	for ; k >= 0; k-- {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
			continue
		}
		if j < 0 {
			nums1[k] = nums1[i]
			i--
			continue
		}
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
	return nums1
}

// @lc code=end
