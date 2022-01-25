package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestLengthOfLongestSubString(t *testing.T) {
	myt := easytest.New(t).F(lengthOfLongestSubstring)
	myt.Run("abcabcbb").Equal(3)
	myt.Run("bbbbb").Equal(1)
	myt.Run("pwwkew").Equal(3)
	myt.Run("").Equal(0)
	myt.Run(" ").Equal(1)
	myt.Run("tmmzuxt").Equal(5)
}

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// last position the byte appeared
	m := make(map[byte]int)
	// start index of last byte
	prev := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if m[s[i]] >= prev {
				prev = m[s[i]] + 1
			}
		}
		m[s[i]] = i
		if i-prev+1 > max {
			max = i - prev + 1
		}
	}
	return max
}

// @lc code=end
func lengthOfLongestSubstring1(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		m := make(map[byte]bool)
		cnt := 0
		for j := i; j < len(s); j++ {
			if m[s[j]] {
				break
			} else {
				m[s[j]] = true
				cnt++
			}
			if cnt > max {
				max = cnt
			}
		}
	}
	return max
}
