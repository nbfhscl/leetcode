package test

import (
	"strconv"
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestCountAndSay(t *testing.T) {
	myt := easytest.New(t)
	myt.F(countAndSay).Run(1).Equal("1")
	myt.Run(2).Equal("11")
	myt.Run(3).Equal("21")
	myt.Run(4).Equal("1211")
	myt.Run(5).Equal("111221")
	myt.Run(6).Equal("312211")
}

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start

func countAndSay(n int) string {
	r := strings.Builder{}
	r.WriteString("1")
	for i := 1; i < n; i++ {
		v := r.String() + " "
		r = strings.Builder{}
		for cnt, i := 1, 0; i < len(v)-1; i, cnt = i+1, cnt+1 {
			if v[i] != v[i+1] {
				r.WriteString(strconv.Itoa(cnt))
				r.WriteByte(v[i])
				cnt = 0
			}
		}
	}
	return r.String()
}

// @lc code=end

func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}
	v := " " + countAndSay(n-1)
	r := ""
	for cnt, d := 0, v[len(v)-1]; len(v) > 0; v, cnt = v[:len(v)-1], cnt+1 {
		if v[len(v)-1] != d {
			r = strconv.Itoa(cnt) + string(d) + r
			d = v[len(v)-1]
			cnt = 0
		}
	}
	return r
}

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	v := countAndSay(n-1) + " "
	r := ""
	for cnt, i := 1, 1; i < len(v); i, cnt = i+1, cnt+1 {
		if v[i] != v[i-1] {
			r += strconv.Itoa(cnt) + string(v[i-1])
			cnt = 0
		}
	}
	return r
}

func countAndSay3(n int) string {
	if n == 1 {
		return "1"
	}
	v := countAndSay(n-1) + " "
	r := ""
	for cnt, i := 1, 0; i < len(v)-1; i, cnt = i+1, cnt+1 {
		if v[i] != v[i+1] {
			r += strconv.Itoa(cnt) + string(v[i])
			cnt = 0
		}
	}
	return r
}

func countAndSay4(n int) string {
	r := "1"
	for i, v := 1, r+" "; i < n; i++ {
		v = r + " "
		r = ""
		for cnt, i := 1, 0; i < len(v)-1; i, cnt = i+1, cnt+1 {
			if v[i] != v[i+1] {
				r += strconv.Itoa(cnt) + string(v[i])
				cnt = 0
			}
		}
	}
	return r
}
