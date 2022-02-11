package test

import (
	"strings"
	"testing"
	"unicode"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestIsPalindrome(t *testing.T) {
	myt := easytest.New(t).F(isPalindrome)
	myt.Run("A man, a plan, a canal: Panama").Equal(true)
	myt.Run("race a car").Equal(false)
}

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start

func isalphanumeric(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isalphanumeric(s[i]) {
			i++
		}
		for i < j && !isalphanumeric(s[j]) {
			j--
		}
		if i < j && toLower(s[i]) != toLower(s[j]) {
			return false
		}
	}
	return true
}

// @lc code=end

func isPalindrome1(s string) bool {
	isalphanumeric := func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		if unicode.IsNumber(r) {
			return r
		}
		return -1
	}
	s = strings.Map(isalphanumeric, s)
	rsb := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		rsb.WriteByte(s[i])
	}
	return s == rsb.String()
}

func isPalindrome2(s string) bool {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !unicode.IsLetter(rs[i]) && !unicode.IsNumber(rs[i]) {
			i++
		}
		for i < j && !unicode.IsLetter(rs[j]) && !unicode.IsNumber(rs[j]) {
			j--
		}
		if unicode.IsUpper(rs[i]) {
			rs[i] = unicode.ToLower(rs[i])
		}
		if unicode.IsUpper(rs[j]) {
			rs[j] = unicode.ToLower(rs[j])
		}
		if i < j && rs[i] != rs[j] {
			return false
		}
	}
	return true
}
