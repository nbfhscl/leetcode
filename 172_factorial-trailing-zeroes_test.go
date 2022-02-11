package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestTrailingZeroes(t *testing.T) {
	myt := easytest.New(t).F(trailingZeroes)
	myt.Input(3).Output(0)
	myt.Input(5).Output(1)
	myt.Input(0).Output(0)
	myt.Input(10).Output(2)
}

/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */

// @lc code=start
func trailingZeroes(n int) int {
	var i, r int = 5, 0
	for n >= i {
		r += n / i
		// each multiple of 5 will contribute
		// another 1 to the trailing zeros
		i *= 5
	}
	return r
}

// @lc code=end
func trailingZeroes1(n int) int {
	num2, num5, num10 := 0, 0, 0
	for i := 1; i <= n; i++ {
		v := i
		for v != 0 && v%10 == 0 {
			num10++
			v /= 10
		}
		for v != 0 && v%5 == 0 {
			num5++
			v /= 5
		}
		for v != 0 && v%2 == 0 {
			num2++
			v /= 2
		}
	}
	trail0 := 0
	if num2 < num5 {
		trail0 = num2
	} else {
		trail0 = num5
	}
	trail0 += num10
	return trail0
}
