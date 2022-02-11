package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestLongestPalindrome(t *testing.T) {
	myt := easytest.New(t).F(longestPalindrome)
	myt.Input("babad").Output("bab")
	myt.Input("cbbd").Output("bb")
	myt.Input("a").Output("a")
	myt.Input("ac").Output("a")
	myt.Input("aaaa").Output("aaaa")
	myt.Input("baaaab").Output("baaaab")
	myt.Input("baaab").Output("baaab")
	myt.Input("baba").Output("bab")
	myt.Input("babc").Output("bab")
	myt.Input("cabc").Output("c")
	myt.Input("cacc").Output("cac")
	myt.Input("baaaabaaaa").Output("aaaabaaaa")
}

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	var mid, left, right, rl, rr int
	for mid < len(s) {
		left, right = mid, mid+1
		// deal with duplicated chars
		for right < len(s) && s[right] == s[mid] {
			right++
		}
		// next possible mid
		mid = right
		// deal with left and right end
		for left > 0 && right < len(s) && s[left-1] == s[right] {
			left--
			right++
		}
		// backup
		if right-left > rr-rl {
			rr, rl = right, left
		}
	}
	return s[rl:rr]
}

// @lc code=end
