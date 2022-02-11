package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMultiply(t *testing.T) {
	myt := easytest.New(t).F(multiply)
	myt.Input("2", "3").Output("6")
	myt.Input("123", "456").Output("56088")

	myt.Input("9", "9").Output("81")
	myt.Input("99999", "99999").Output("9999800001")
}

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	resd := make([]int, len(num1)+len(num2))
	for i := range num1 {
		for j := range num2 {
			resd[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}

	for i := len(resd) - 1; i > 0; i-- {
		if resd[i] > 9 {
			resd[i-1] += resd[i] / 10
			resd[i] = resd[i] % 10
		}
	}
	if resd[0] == 0 {
		resd = resd[1:]
	}

	res := make([]byte, len(resd))
	for i := range resd {
		res[i] = byte(resd[i]) + '0'
	}
	return string(res)
}

// @lc code=end

//      4321
// *     321
// ---------
// +    4321
// +   8642
// + 12963
// = 1387041
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	mid := make([][]byte, len(num2))
	for i := range num2 {
		mid[i] = make([]byte, len(num1)+len(num2))
		for j := range num1 {
			mid[i][j+i] += (num2[len(num2)-1-i] - '0') * (num1[len(num1)-1-j] - '0')
			if mid[i][j+i] > 9 {
				mid[i][j+i+1] = mid[i][j+i] / 10
				mid[i][j+i] = mid[i][j+i] % 10
			}
		}
	}
	resd := make([]byte, len(num1)+len(num2))
	for i := range resd {
		for j := range mid {
			resd[i] += mid[j][i]
			if resd[i] > 9 {
				resd[i+1] += resd[i] / 10 // 注意这里要累加
				resd[i] = resd[i] % 10
			}
		}
	}
	var i int
	for i = len(resd) - 1; i >= 0; i-- {
		if resd[i] != 0 {
			break
		}
	}
	res := strings.Builder{}
	for ; i >= 0; i-- {
		res.WriteByte(resd[i] + '0')
	}
	return res.String()
}
