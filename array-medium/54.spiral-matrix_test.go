package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSpiralOrder(t *testing.T) {
	myt := easytest.New(t).F(spiralOrder)
	myt.Input([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}).Output([]int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	myt.Input([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}).Output([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}
	jmax, imax := len(matrix[0])-1, len(matrix)-1
	jmin, imin := 0, 1
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for dir := 1; ; dir = -dir {
		if dir == 1 {
			for j := jmin; j <= jmax; j++ {
				res = append(res, matrix[imin-1][j])
			}
			for i := imin; i <= imax; i++ {
				res = append(res, matrix[i][jmax])
			}
			imax, jmax = imax-1, jmax-1
		} else {
			for j := jmax; j >= jmin; j-- {
				res = append(res, matrix[imax+1][j])
			}
			for i := imax; i >= imin; i-- {
				res = append(res, matrix[i][jmin])
			}
			imin, jmin = imin+1, jmin+1
		}
		if len(res) == cap(res) {
			return res
		}
	}
}

// @lc code=end
