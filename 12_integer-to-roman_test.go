package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIntToRoman(t *testing.T) {
	myt := easytest.New(t).F(intToRoman)
	myt.Input(4).Output("IV")
	myt.Input(9).Output("IX")
	myt.Input(90).Output("XC")
	myt.Input(40).Output("XL")
	myt.Input(400).Output("CD")
	myt.Input(900).Output("CM")
	myt.Input(3).Output("III")
	myt.Input(58).Output("LVIII")
	myt.Input(1994).Output("MCMXCIV")
}

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 */

// @lc code=start
func intToRoman(num int) string {
	integers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romannum := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	bs := strings.Builder{}
	i := 0
	for num > 0 {
		for num < integers[i] {
			i++
		}
		bs.WriteString(romannum[i])
		num -= integers[i]
	}
	return bs.String()
}

// @lc code=end
func intToRoman1(num int) string {
	bs := strings.Builder{}
	for num > 0 {
		switch {
		case num >= 1000:
			bs.WriteString("M")
			num -= 1000
		case num >= 900:
			bs.WriteString("CM")
			num -= 900
		case num >= 500:
			bs.WriteString("D")
			num -= 500
		case num >= 400:
			bs.WriteString("CD")
			num -= 400
		case num >= 100:
			bs.WriteString("C")
			num -= 100
		case num >= 90:
			bs.WriteString("XC")
			num -= 90
		case num >= 50:
			bs.WriteString("L")
			num -= 50
		case num >= 40:
			bs.WriteString("XL")
			num -= 40
		case num >= 10:
			bs.WriteString("X")
			num -= 10
		case num >= 9:
			bs.WriteString("IX")
			num -= 9
		case num >= 5:
			bs.WriteString("V")
			num -= 5
		case num >= 4:
			bs.WriteString("IV")
			num -= 4
		case num >= 1:
			bs.WriteString("I")
			num -= 1
		}
	}
	return bs.String()
}
