package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestDivide(t *testing.T) {
	myt := easytest.New(t).F(divide)
	myt.Input(10, 3).Output(3)
	myt.Input(0, 1).Output(0)
	myt.Input(1, 1).Output(1)
	myt.Input(1000000, 10).Output(100000)
	myt.Input(2147483647, 2).Output(1073741823)

	myt.Input(7, -3).Output(-2)
	myt.Input(3, -1).Output(-3)
	myt.Input(10000, -10).Output(-1000)

	myt.Input(math.MaxInt32, 1).Output(math.MaxInt32)
	myt.Input(math.MaxInt32, -1).Output(math.MinInt32 + 1)

	// corner cases
	myt.Input(math.MinInt32, 1).Output(math.MinInt32)
	myt.Input(-2147483648, 2).Output(-1073741824)
	myt.Input(math.MinInt32, -1).Output(math.MaxInt32)
	myt.Input(-2147483648, -2).Output(1073741824)
	myt.Input(-2147483648, -2147483648).Output(1)
	myt.Input(-1, -2147483648).Output(0)
	myt.Input(-2, -2147483648).Output(0)
	myt.Input(1, -2147483648).Output(0)
	myt.Input(2, -2147483648).Output(0)
}

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start

func divide2(dividend int, divisor int) int {
	var a, b int32 = int32(dividend), int32(divisor)
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	var neg int32 = 1
	if a > 0 {
		neg = -neg
		a = -a
	}
	if b > 0 {
		neg = -neg
		b = -b
	}
	var res int32

	return int(neg * res)
}

// @lc code=end
// 20 / 3
// 20 - 3 > 0   0 17
// 20 - 3*2 > 0 2 14
// 20 - 3*4 > 0 4 8
// 20 - 3*8 < 0 8 -4
// 8 - 3 > 0    0 5
// 8 - 3*2 > 0  2 2
// 8 - 3*4 < 0  4 -4
func divide(dividend int, divisor int) int {
	var a, b int32 = int32(dividend), int32(divisor)
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	var neg int32 = 1
	if a < 0 {
		neg = -neg
		a = -a // math.MinInt32 == -math.MinInt32
	}
	if b < 0 {
		neg = -neg
		b = -b
	}
	var res, x int32
	for a-b >= 0 { // math.MinInt32 - positive >= 0 || math.MinInt32 - math.MinInt32 == 0
		for x = 0; a-b<<x<<1 >= 0; x++ {
		}
		res += 1 << x
		a -= b << x
	}
	return int(neg * res)
}
