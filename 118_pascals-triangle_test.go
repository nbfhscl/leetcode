package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGenerate(t *testing.T) {
	myt := easytest.New(t).F(generate)
	myt.Input(1).Output([][]int{{1}})
	myt.Input(2).Output([][]int{{1}, {1, 1}})
	myt.Input(3).Output([][]int{{1}, {1, 1}, {1, 2, 1}})
	myt.Input(4).Output([][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}})
	myt.Input(5).Output([][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
}

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	r := make([][]int, numRows)
	for i := range r {
		r[i] = make([]int, i+1)
		r[i][0], r[i][i] = 1, 1
		for j := 1; j < i; j++ {
			r[i][j] = r[i-1][j-1] + r[i-1][j]
		}
	}
	return r
}

// @lc code=end
