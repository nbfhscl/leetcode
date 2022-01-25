package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestConvertToTitle(t *testing.T) {
	myt := easytest.New(t).F(convertToTitle)
	myt.Input(1).Output("A")
	myt.Input(2).Output("B")
	myt.Input(3).Output("C")
	myt.Input(26).Output("Z")
	myt.Input(27).Output("AA")
	myt.Input(701).Output("ZY")

}

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
// 0 - 25
// 1 - 26
// 1 -> A -> 'A' +0
// 26 -> Z -> 'A' +25
// 27 -> AA -> 'A' +0
func convertToTitle(n int) string {
	r := make([]byte, 0, 3)
	for n > 0 {
		n--
		r = append(r, 'A'+byte(n%26))
		n /= 26
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// @lc code=end
