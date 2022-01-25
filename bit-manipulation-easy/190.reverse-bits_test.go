package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestReverseBits(t *testing.T) {
	myt := easytest.New(t).F(reverseBits)
	myt.Run(uint32(0b00000010100101000001111010011100)).Equal(uint32(964176192))
	myt.Run(uint32(0b11111111111111111111111111111101)).Equal(uint32(3221225471))
}

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var r uint32
	for i := 31; i >= 0; i-- {
		r += (num & 0x01) << i
		num >>= 1
	}
	return r
}

// @lc code=end
