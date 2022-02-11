package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCombine(t *testing.T) {
	myt := easytest.New(t).F(combine)
	myt.Input(4, 2).Output([][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}})
	myt.Input(1, 1).Output([][]int{{1}})
	myt.Input(5, 1).Output([][]int{{1}, {2}, {3}, {4}, {5}})
}

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	btCombine(&res, []int{}, n, k)
	return res
}

func btCombine(res *[][]int, ans []int, n, k int) {
	if len(ans) == k {
		*res = append(*res, append([]int{}, ans...))
	} else {
		for i := n; i >= k-len(ans); i-- {
			btCombine(res, append(ans, i), i-1, k)
		}
	}
}

// @lc code=end
