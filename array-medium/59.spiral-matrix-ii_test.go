package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGenerateMatrix(t *testing.T) {
	myt := easytest.New(t).F(generateMatrix)
	// myt.Input(3).Output([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}})
	// myt.Input(1).Output([][]int{{1}})
	myt.Input(2).Output([][]int{{1, 2}, {4, 3}})
}

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	jmax, imax := n-1, n-1
	jmin, imin := 0, 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	value := 0
	for dir := 1; ; dir = -dir {
		if dir == 1 {
			for j := jmin; j <= jmax; j++ {
				value++
				res[imin-1][j] = value
			}
			for i := imin; i <= imax; i++ {
				value++
				res[i][jmax] = value
			}
			imax, jmax = imax-1, jmax-1
		} else {
			for j := jmax; j >= jmin; j-- {
				value++
				res[imax+1][j] = value
			}
			for i := imax; i >= imin; i-- {
				value++
				res[i][jmin] = value
			}
			imin, jmin = imin+1, jmin+1
		}
		if value == n*n {
			return res
		}
	}
}

// @lc code=end
