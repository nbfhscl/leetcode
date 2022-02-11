package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMinPathSum(t *testing.T) {
	myt := easytest.New(t).F(minPathSum)
	myt.Input([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}).Output(7)
	myt.Input([][]int{{1, 2, 3}, {4, 5, 6}}).Output(12)
}

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	cache[0][0] = grid[0][0]
	return dpMinPathSum(grid, cache, m-1, n-1)
}

func dpMinPathSum(grid, cache [][]int, i, j int) int {
	if cache[i][j] != -1 {
		return cache[i][j]
	}
	left, top := -1, -1
	if i > 0 {
		top = dpMinPathSum(grid, cache, i-1, j)
	}
	if j > 0 {
		left = dpMinPathSum(grid, cache, i, j-1)
	}
	if top == -1 {
		cache[i][j] = left
	} else if left == -1 {
		cache[i][j] = top
	} else {
		if top > left {
			cache[i][j] = left
		} else {
			cache[i][j] = top
		}
	}
	cache[i][j] += grid[i][j]
	return cache[i][j]
}

// @lc code=end
