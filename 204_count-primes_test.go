package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCountPrimes(t *testing.T) {
	myt := easytest.New(t).F(countPrimes)
	myt.Input(10).Output(4)
	myt.Input(2).Output(0)
	myt.Input(100).Output(25)
}

/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	if n == 3 {
		return 1
	}
	cnt := 1
outerfor:
	for last := 3; last < n; last += 2 {
		for i := 3; i <= int(math.Sqrt(float64(last))); i += 2 {
			if last%i == 0 {
				continue outerfor
			}
		}
		cnt++
	}
	return cnt
}

// @lc code=end
