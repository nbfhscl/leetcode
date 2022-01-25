package test

import (
	"math"
	"strconv"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMyAtoi(t *testing.T) {
	myt := easytest.New(t).F(myAtoi)
	myt.Input("").Output(0)
	myt.Input(" ").Output(0)
	myt.Input("+").Output(0)
	myt.Input("-").Output(0)
	myt.Input("+0").Output(0)
	myt.Input("-0").Output(0)
	myt.Input("+1").Output(1)
	myt.Input("-1").Output(-1)
	myt.Input("1").Output(1)
	// 10 digits out of range int32
	myt.Input("9999999999").Output(math.MaxInt32)
	// 11 digits out of range int32
	myt.Input("99999999999").Output(math.MaxInt32)
	// 19 digits out of range int64
	myt.Input("9999999999999999999").Output(math.MaxInt32)
	// 10 digits out of range int32
	myt.Input("-99999999999").Output(math.MinInt32)
	// 11 digits out of range int32
	myt.Input("-999999999999").Output(math.MinInt32)
	// 19 digits out of range int32
	myt.Input("-9999999999999999999").Output(math.MinInt32)
	// 15 digits in range int32
	myt.Input("-000000009999999").Output(-9999999)
	myt.Input("    +1").Output(1)
	myt.Input("   +0001").Output(1)
	myt.Input("   +01001").Output(1001)
	myt.Input("   + 01001").Output(0)
	myt.Input("   +010 01").Output(10)
	myt.Input(strconv.Itoa(math.MaxInt32)).Output(math.MaxInt32)
	myt.Input(strconv.Itoa(math.MinInt32)).Output(math.MinInt32)
}

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	var r int
	var sign int = 1
	i := 0
	// Read in and ignore leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i >= len(s) {
		return 0
	}
	// Check sign
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	default:
	}
	for i < len(s) && s[i] == '0' {
		i++
	}
	// Read in digits and convert it to integer
	for j := 1; i < len(s); i, j = i+1, j+1 {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		r *= 10
		// apply sign
		r += sign * int(s[i]-'0')
		if j > 10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		} else if j == 10 {
			// Check result in range
			if r >= math.MaxInt32 {
				r = math.MaxInt32
				break
			}
			if r <= math.MinInt32 {
				r = math.MinInt32
				break
			}
		}
	}
	return r
}

// @lc code=end
