package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestExist(t *testing.T) {
	myt := easytest.New(t).F(exist)
	myt.Input([][]byte{{'A', 'B', 'C'}}, "ABC").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "DEF").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABE").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABED").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABEDA").Output(false)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABEF").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABEFC").Output(true)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "ABEFCB").Output(false)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "J").Output(false)
	myt.Input([][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}}, "abc").Output(false)
	myt.Input([][]byte{{'A'}}, "A").Output(true)
	myt.Input([][]byte{{'A'}}, "a").Output(false)

	//
	myt.Input([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB").Output(true)
}

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	bWord := []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == bWord[0] {
				if btExist(board, bWord, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func btExist(board [][]byte, bWord []byte, m, n int) bool {
	if len(bWord) == 1 {
		return true
	}
	board[m][n] = 0
	if m > 0 && board[m-1][n] == bWord[1] {
		if btExist(board, bWord[1:], m-1, n) {
			return true
		}
	}
	if n > 0 && board[m][n-1] == bWord[1] {
		if btExist(board, bWord[1:], m, n-1) {
			return true
		}
	}
	if m < len(board)-1 && board[m+1][n] == bWord[1] {
		if btExist(board, bWord[1:], m+1, n) {
			return true
		}
	}
	if n < len(board[0])-1 && board[m][n+1] == bWord[1] {
		if btExist(board, bWord[1:], m, n+1) {
			return true
		}
	}
	board[m][n] = bWord[0]
	return false
}

// @lc code=end
