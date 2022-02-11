package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestLengthOfLastWord(t *testing.T) {
	myt := easytest.New(t)
	myt.F(lengthOfLastWord).Run("Hello World").Equal(5)
	myt.F(lengthOfLastWord).Run("Hello World ").Equal(5)
	myt.F(lengthOfLastWord).Run("H").Equal(1)
	myt.F(lengthOfLastWord).Run("    ").Equal(0)
	myt.F(lengthOfLastWord).Run("  Hello World").Equal(5)
}

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ss := strings.TrimRight(s, " ")
	return len(ss) - 1 - strings.LastIndex(ss, " ")
}

// @lc code=end

func lengthOfLastWord1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var r int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if r == 0 {
				continue
			}
			return r
		}
		r++
	}
	return r
}
