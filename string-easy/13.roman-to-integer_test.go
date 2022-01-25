package test

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tcs := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "III",
			input:  "III",
			output: 3,
		},
		{
			name:   "XIV",
			input:  "XIV",
			output: 14,
		},
		{
			name:   "IV",
			input:  "IV",
			output: 4,
		},
		{
			name:   "IX",
			input:  "IX",
			output: 9,
		},
		{
			name:   "XL",
			input:  "XL",
			output: 40,
		},
		{
			name:   "XC",
			input:  "XC",
			output: 90,
		},
		{
			name:   "CD",
			input:  "CD",
			output: 400,
		},
		{
			name:   "CM",
			input:  "CM",
			output: 900,
		},
		{
			name:   "MCMXCIV",
			input:  "MCMXCIV",
			output: 1994,
		},
		{
			name:   "LVIII",
			input:  "LVIII",
			output: 58,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := romanToInt(tc.input)
			if tc.output != output {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	var v int
	length := len(s)
	for i, j := 0, 0; i < length; i++ {
		if i == length-1 {
			j = i + 1
		} else {
			j = i + 2
		}
		switch s[i:j] {
		case "IV":
			v += 4
			i++
		case "IX":
			v += 9
			i++
		case "XL":
			v += 40
			i++
		case "XC":
			v += 90
			i++
		case "CD":
			v += 400
			i++
		case "CM":
			v += 900
			i++
		default:
			switch s[i : i+1] {
			case "I":
				v += 1
			case "V":
				v += 5
			case "X":
				v += 10
			case "L":
				v += 50
			case "C":
				v += 100
			case "D":
				v += 500
			case "M":
				v += 1000
			default:
				return v
			}
		}
	}
	return v
}

// @lc code=end
