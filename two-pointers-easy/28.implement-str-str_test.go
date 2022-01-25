package test

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tcs := []struct {
		name     string
		haystack string
		needle   string
		output   int
	}{
		{
			name:     "haystack is empty while needle isnt",
			haystack: "",
			needle:   "fjjf",
			output:   -1,
		},
		{
			name:     "haystak is empty and needle too",
			haystack: "",
			needle:   "",
			output:   0,
		},
		{
			name:     "example1",
			haystack: "hello",
			needle:   "ll",
			output:   2,
		},
		{
			name:     "example2",
			haystack: "aaaaa",
			needle:   "bba",
			output:   -1,
		},
		{
			name:     "needle longer than haystack",
			haystack: "abcdefg",
			needle:   "abcdefgh",
			output:   -1,
		},
		{
			name:     "haystack overhead",
			haystack: "fhoeijfabcde",
			needle:   "abcdef",
			output:   -1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := strStr(tc.haystack, tc.needle)
			if tc.output != output {
				t.Errorf("input: %v, output: %v, expected: %v", tc.haystack+"/"+tc.needle, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}

	for i, v := range []byte(haystack)[:len(haystack)-len(needle)+1] {
		if v == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
