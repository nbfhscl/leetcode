package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

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
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}

// @lc code=end
