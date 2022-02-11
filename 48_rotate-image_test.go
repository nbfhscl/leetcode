package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestRotate(t *testing.T) {
	myt := easytest.New(t).F(rotate)
	myt.Input([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}).Output([][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
}

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

// @lc code=end
