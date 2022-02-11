package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMyPow(t *testing.T) {
	myt := easytest.New(t).F(myPow)
	myt.Input(2.00000, 10).Output(1024.00000)
	myt.Input(2.10000, 3).Output(9.26100)
	myt.Input(2.00000, -2).Output(0.25000)
}

/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / (x * myPow(x, -(n+1)))
	}
	res := 1.0
	for n > 1 {
		if n%2 == 1 {
			res *= x
		}
		x *= x
		n /= 2
	}
	return res * x
}

// @lc code=end
