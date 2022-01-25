package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIsValidSudoku(t *testing.T) {
	myt := easytest.New(t).F(isValidSudoku)
	myt.Input([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}).Output(true)
	myt.Input([][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}).Output(false)
}

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	block := make([][]int, 10)
	for i := range block {
		block[i] = make([]int, 10)
	}
	for i := 0; i < 9; i++ {
		row := make([]int, 10)
		col := make([]int, 10)
		for j := 0; j < 9; j++ {
			v1, v2 := board[i][j], board[j][i]
			if v1 != '.' {
				v1 -= '0'
				if row[v1] == 1 {
					return false
				}
				row[v1] = 1
				hash := 3*(i/3) + (j / 3)
				if block[hash][v1] == 1 {
					return false
				}
				block[hash][v1] = 1
			}
			if v2 != '.' {
				v2 -= '0'
				if col[v2] == 1 {
					return false
				}
				col[v2] = 1
			}
		}
	}
	return true
}

// @lc code=end
