package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGrayCode(t *testing.T) {
	myt := easytest.New(t).F(grayCode)
	myt.Input(2).Output([]int{0, 1, 3, 2})
	myt.Input(1).Output([]int{0, 1})

	myt.Input(3).Output([]int{0, 1, 3, 2, 6, 7, 5, 4})
}

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
func grayCode(n int) []int {
	if n == 1 {
		return []int{0, 1}
	}

	return append(grayCode(n-1), reverseOr(grayCode(n-1), int(math.Pow(2, float64(n-1))))...)
}

func reverseOr(po []int, n int) []int {
	for i, j := 0, len(po)-1; i < j; i, j = i+1, j-1 {
		po[i], po[j] = n|po[j], n|po[i]
	}
	return po
}

// @lc code=end

// 0, 1
// 0, 1, 3, 2
// 0, 1, 3, 2, 6, 7, 5, 4

// 0, 1
// 00, 01, 11, 10
// 000, 001, 011, 010, 110, 111, 101, 100
