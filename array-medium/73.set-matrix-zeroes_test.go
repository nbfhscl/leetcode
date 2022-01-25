package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSetZeroes(t *testing.T) {
	myt := easytest.New(t).F(setZeroes)
	myt.Input([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}).Output([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
	myt.Input([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}).Output([][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}})
}

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				markRowAndCol(matrix, i, j)
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == math.MaxInt32+1 {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func markRowAndCol(matrix [][]int, m, n int) {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][n] != 0 {
			matrix[i][n] = math.MaxInt32 + 1
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[m][j] != 0 {
			matrix[m][j] = math.MaxInt32 + 1
		}
	}
}

// @lc code=end
