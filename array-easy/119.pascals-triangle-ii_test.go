package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGetRow(t *testing.T) {
	myt := easytest.New(t).F(getRow)
	myt.Input(0).Output([]int{1})
	myt.Input(1).Output([]int{1, 1})
	myt.Input(2).Output([]int{1, 2, 1})
	myt.Input(3).Output([]int{1, 3, 3, 1})
	myt.Input(4).Output([]int{1, 4, 6, 4, 1})
}

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	var temp int
	for i := range r {
		temp = 1
		r[0], r[i] = 1, 1
		for j := 1; j < i; j++ {
			r[j], temp = temp+r[j], r[j]
		}
	}
	return r
}

// @lc code=end
