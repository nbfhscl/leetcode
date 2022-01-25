package test

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tcs := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "positive sign",
			input:  +121,
			output: true,
		},
		{
			name:   "negative sign",
			input:  -121,
			output: false,
		},
		{
			name:   "101101",
			input:  101101,
			output: true,
		},
		{
			name:   "10",
			input:  10,
			output: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := isPalindrome(tc.input)
			if tc.output != output {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var y int
	xx := x
	for xx != 0 {
		y = y*10 + xx%10
		xx = xx / 10
	}
	return x == y
}

// @lc code=end

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	var digits [10]int8
	var i int8 = -1
	for x != 0 {
		i++
		digits[i] = int8(x % 10)
		x = x / 10
	}
	var j int8
	for i > j {
		if digits[i] != digits[j] {
			return false
		}
		j++
		i--
	}
	return true
}

func isPalindrome3(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	var y int
	for x > y {
		y = y*10 + x%10
		x = x / 10
		if x == y || x/10 == y {
			return true
		}
	}
	return false
}
