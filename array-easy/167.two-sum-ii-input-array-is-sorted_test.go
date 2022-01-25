package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestTwoSum(t *testing.T) {
	myt := easytest.New(t).F(twoSum)
	myt.Input([]int{2, 7, 11, 15}, 9).Output([]int{1, 2})
	myt.Input([]int{2, 3, 4}, 6).Output([]int{1, 3})
	myt.Input([]int{-1, 0}, -1).Output([]int{1, 2})
}

/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for i, v := range numbers {
		if _, ok := m[v]; ok {
			return []int{m[v], i + 1}
		} else {
			m[target-v] = i + 1
		}
	}
	return []int{}
}

// @lc code=end
