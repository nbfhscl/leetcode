package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestUniquePaths(t *testing.T) {
	myt := easytest.New(t).F(uniquePaths)
	myt.Input(3, 7).Output(28)
	myt.Input(3, 2).Output(3)
	myt.Input(7, 3).Output(28)
	myt.Input(3, 3).Output(6)
}

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start

// node[i][j] = node[i-1][j] + node[i][j-1]
func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	return dpUniquePaths(cache, m-1, n-1)
}

func dpUniquePaths(cache [][]int, m, n int) int {
	if cache[m][n] != 0 {
		return cache[m][n]
	}
	if m == 0 || n == 0 {
		return 1
	}
	cache[m][n] = dpUniquePaths(cache, m, n-1) + dpUniquePaths(cache, m-1, n)
	return cache[m][n]
}

// @lc code=end
