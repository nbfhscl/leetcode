package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMySqrt(t *testing.T) {
	myt := easytest.New(t)
	myt.F(mySqrt)
	myt.Input(0).Output(0)
	myt.Input(1).Output(1)
	myt.Input(4).Output(2)
	myt.Input(8).Output(2)
	myt.Input(9).Output(3)
	myt.Input(100).Output(10)
}

func TestSqrtNewton(t *testing.T) {
	myt := easytest.New(t)
	myt.F(SqrtNewton).Input(0.0).Output(0)
	myt.Input(math.Inf(1)).Output(math.MaxInt64)
	myt.Input(math.Inf(-1)).Output(math.MinInt64)
	myt.Input(1.0).Output(1)
	myt.Input(4.0).Output(2)
	myt.Input(8.0).Output(2)
	myt.Input(9.0).Output(3)
	myt.Input(100.0).Output(10)
}

func TestSqrtDichotomy(t *testing.T) {
	myt := easytest.New(t)
	myt.F(SqrtDichotomy).Input(0.0).Output(0)
	// myt.Input(math.Inf(1)).Output(math.MaxInt64)
	// myt.Input(math.Inf(-1)).Output(math.MinInt64)
	myt.Input(1.0).Output(1)
	myt.Input(4.0).Output(2)
	myt.Input(8.0).Output(2)
	myt.Input(9.0).Output(3)
	myt.Input(100.0).Output(10)
}

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func SqrtNewton(x float64) float64 {
	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	Precision := 0.0001
	nt := x
	for nt*nt-Precision > x || nt*nt+Precision < x {
		nt = 0.5 * (nt + x/nt)
	}
	return nt
}
func mySqrt(x int) int {
	return int(SqrtNewton(float64(x)))
	// return int(math.Sqrt(float64(x)))
}

// @lc code=end

func SqrtDichotomy(x float64) float64 {
	Precision := 0.0001
	min := 0.0
	max := x
	for {
		mid := 0.5 * (min + max)
		if mid*mid-Precision > x {
			max = mid
			continue
		}
		if mid*mid+Precision < x {
			min = mid
			continue
		}
		return mid
	}
}
