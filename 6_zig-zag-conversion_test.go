package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestConvert(t *testing.T) {
	myt := easytest.New(t).F(convert)
	myt.Input("PAYPALISHIRING", 3).Output("PAHNAPLSIIGYIR")
	myt.Input("PAYPALISHIRING", 4).Output("PINALSIGYAHRPI")
	myt.Input("A", 1).Output("A")
}

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	sb := strings.Builder{}
	for row := 1; row <= numRows; row++ {
		for i, index := 1, row-1; index < len(s); i, index = i+1, row-1+2*i*(numRows-1) {
			sb.WriteByte(s[index])
			if row > 1 && row < numRows && index+2*(numRows-row) < len(s) {
				sb.WriteByte(s[index+2*(numRows-row)])
			}
		}

	}
	return sb.String()
}

// @lc code=end

// 2(n-1)-2 = 2n-4 = 4
// 2(n-2)-2 = 2n-6 = 2
// 2(n-row+1)-2= 2(n-row)

// 2(n-1) = 2n-2	i(2n-2)
// 2(n-2) = 2n-4	1+i(2n-4)
// 2(n-3) = 2n-6	2+i(2n-6)
// 2(n-row)		row-1+2i(n-row)
