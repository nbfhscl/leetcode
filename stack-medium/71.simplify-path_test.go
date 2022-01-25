package test

import (
	"strings"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestSimplifyPath(t *testing.T) {
	myt := easytest.New(t).F(simplifyPath)
	myt.Input("/home/").Output("/home")
	myt.Input("/../").Output("/")
	myt.Input("/home//foo/").Output("/home/foo")
	myt.Input("/a/./b/../../c/").Output("/c")
}

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start

func simplifyPath(path string) string {
	stack := make([]string, 0)
	strs := strings.Split(path, "/")
	for i := range strs {
		if strs[i] == "" || strs[i] == "." {
			continue
		} else if strs[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, strs[i])
		}
	}
	str := ""
	for i := range stack {
		str += "/" + stack[i]
	}
	if str == "" {
		str = "/"
	}
	return str
}

// @lc code=end
