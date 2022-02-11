package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestAddBinary(t *testing.T) {
	myt := easytest.New(t)
	myt.F(addBinary)
	myt.Input("1111", "1111").Output("11110")
	myt.Input("101011", "101110").Output("1011001")
	myt.Input("1010", "1011").Output("10101")
	myt.Input("101111", "10").Output("110001")
}

/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var c byte
	sb := &strings.Builder{}
	for i, j := len(a)-1, len(b)-1; i >= 0; i, j = i-1, j-1 {
		var nb byte
		if j >= 0 {
			nb = b[j] - '0'
		}
		na := a[i] - '0'
		sb.WriteByte(na ^ nb ^ c + '0')
		c = na&nb | c&(na|nb)
	}
	if c == 1 {
		sb.WriteByte('1')
	}
	r := []byte(sb.String())
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// @lc code=end

func addBinary1(a string, b string) string {
	var r []byte
	if len(a) > len(b) {
		r = make([]byte, len(a)+1)
	} else {
		r = make([]byte, len(b)+1)
	}
	for i, j, c := len(r)-1, 0, byte(0); i >= 0; i, j = i-1, j+1 {
		var na, nb byte
		if len(a)-1-j >= 0 {
			na = a[len(a)-1-j] - '0'
		}
		if len(b)-1-j >= 0 {
			nb = b[len(b)-1-j] - '0'
		}
		r[i] = na ^ nb ^ c + '0'
		c = na&nb | c&(na|nb)
	}
	if r[0] == '0' || r[0] == 0 {
		return string(r[1:])
	}
	return string(r)
}
