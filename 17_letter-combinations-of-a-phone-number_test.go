package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestLetterCombinations(t *testing.T) {
	myt := easytest.New(t).F(letterCombinations)
	myt.Input("23").Output([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
	myt.Input("").Output([]string{})
	myt.Input("2").Output([]string{"a", "b", "c"})
}

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

var m map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	backtracking(digits, "", &res)
	return res
}

func backtracking(str string, s string, res *[]string) {
	if str == "" {
		if s != "" {
			*res = append(*res, s)
		}
		return
	}
	for _, v := range m[str[0]] {
		s += v
		backtracking(str[1:], s, res)
		s = s[:len(s)-1]
	}
}

// @lc code=end
