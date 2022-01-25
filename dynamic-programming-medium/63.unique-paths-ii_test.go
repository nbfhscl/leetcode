package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	myt := easytest.New(t).F(uniquePathsWithObstacles)
	myt.Input([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}).Output(2)
	myt.Input([][]int{{0, 1}, {0, 0}}).Output(1)

	//
	myt.Input([][]int{{1, 0}}).Output(0)
	myt.Input([][]int{{0, 0}, {1, 1}, {0, 0}}).Output(0)
}

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	return dpUniquePathsWithObstacles(obstacleGrid, m-1, n-1)
}

func dpUniquePathsWithObstacles(obstacleGrid [][]int, m, n int) int {
	if obstacleGrid[m][n] == -1 {
		return 0
	}
	if obstacleGrid[m][n] != 0 {
		return obstacleGrid[m][n]
	}
	if m == 0 && n == 0 {
		return 1
	}
	left, top := 0, 0
	if m > 0 {
		top = dpUniquePathsWithObstacles(obstacleGrid, m-1, n)
	}
	if n > 0 {
		left = dpUniquePathsWithObstacles(obstacleGrid, m, n-1)
	}
	obstacleGrid[m][n] = left + top
	return obstacleGrid[m][n]
}

// @lc code=end
