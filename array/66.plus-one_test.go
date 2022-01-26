package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestPlusOne(t *testing.T) {
	myt := easytest.New(t).F(plusOne)
	myt.Input([]int{1, 2, 3}).Output([]int{1, 2, 4})
	myt.Input([]int{4, 3, 2, 1}).Output([]int{4, 3, 2, 2})
	myt.Input([]int{0}).Output([]int{1})
}

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += c
		c = digits[i] / 10
		if c == 0 {
			return digits
		}
		digits[i] %= 10
	}
	if c == 0 {
		return digits
	}
	arr := make([]int, 1)
	i := 0
	for c != 0 {
		arr[i] = c % 10
		c = c / 10
		i++
	}
	arr2 := make([]int, len(arr))
	for i := range arr2 {
		arr2[i] = arr[len(arr)-1-i]
	}
	return append(arr2, digits...)
}

// @lc code=end
