package test

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tcs := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "123",
			input:  123,
			output: 321,
		},
		{
			name:   "sign",
			input:  -123,
			output: -321,
		},
		{
			name:   "trailing_zero",
			input:  123000,
			output: 321,
		},
		{
			name:   "1534236469",
			input:  1534236469,
			output: 0,
		},
		{
			name:   "overflow_positive",
			input:  2147483647,
			output: 0,
		},
		{
			name:   "overflow_negative",
			input:  -2147483648,
			output: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := reverse(tc.input)
			if tc.output != output {
				t.Errorf("intput: %v, output: %v, expected %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse int32eger
 */

// @lc code=start
func reverse(x int) int {
	var rev int
	var digt int
	var sign bool
	var i32dvd10 int = 214748364
	if x < 0 {
		sign = true
		x = -x
		if x < 0 {
			return 0
		}
	}
	for x != 0 {
		if rev > i32dvd10 {
			return 0
		}
		digt = x % 10
		if rev == i32dvd10 && (sign && digt > 8 || !sign && digt > 7) {
			return 0
		}
		rev = rev*10 + digt
		x = x / 10
	}
	if sign {
		rev = -rev
	}
	return rev
}

// @lc code=end
