package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestTitleToNumber(t *testing.T) {
	myt := easytest.New(t).F(titleToNumber)
	myt.Input("A").Output(1)
	myt.Input("AB").Output(28)
	myt.Input("ZY").Output(701)

}

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(s string) int {
	lt := []byte(s)
	number := 0
	for i := 0; i < len(lt); i++ {
		number *= 26
		number += int(lt[i] - 'A' + 1)
	}
	return number
}

// @lc code=end
