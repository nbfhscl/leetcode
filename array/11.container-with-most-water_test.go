package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMaxArea(t *testing.T) {
	myt := easytest.New(t).F(maxArea)
	myt.Input([]int{1, 1}).Output(1)
	myt.Input([]int{4, 3, 2, 1, 4}).Output(16)
	myt.Input([]int{1, 2, 1}).Output(2)
	myt.Input([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}).Output(49)
}

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1
	for left < right {
		if curr := min(height[left], height[right]) * (right - left); curr > max {
			max = curr
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

func maxArea1(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			curr := min(height[i], height[j]) * (j - i)
			if curr > max {
				max = curr
			}
		}
	}
	return max
}
