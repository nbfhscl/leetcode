package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestHammingWeight(t *testing.T) {
	myt := easytest.New(t).F(hammingWeight)
	myt.Run(uint32(0b00000000000000000000000000001011)).Equal(3)
	myt.Run(uint32(0b00000000000000000000000010000000)).Equal(1)
	myt.Run(uint32(0b11111111111111111111111111111101)).Equal(31)
}

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		cnt += int(num & 0x01)
		num >>= 1
	}
	return cnt
}

// @lc code=end
