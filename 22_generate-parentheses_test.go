package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestGenerateParenthesis(t *testing.T) {
	myt := easytest.New(t).F(generateParenthesis)
	myt.Input(3).Output([]string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	myt.Input(1).Output([]string{"()"})
	myt.Input(2).Output([]string{"(())", "()()"})
	myt.Input(4).Output([]string{"(((())))",
		"((()()))", "((())())", "((()))()",
		"(()(()))", "(()()())", "(()())()", "(())(())", "(())()()",
		"()((()))", "()(()())", "()(())()", "()()(())", "()()()()"})
}

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtracking_parentheses(n, n, "", &res)
	return res
}

func backtracking_parentheses(l, r int, s string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, s)
		return
	}

	if l > 0 {
		s += "("
		backtracking_parentheses(l-1, r, s, res)
		s = s[:len(s)-1]
	}
	if r > 0 && r > l {
		s += ")"
		backtracking_parentheses(l, r-1, s, res)
		s = s[:len(s)-1]
	}
}

// @lc code=end
