package test

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tcs := []struct {
		name   string
		input  []string
		output string
	}{
		{
			name:   "fl",
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			name:   "no common prefix",
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
		{
			name:   "empty input",
			input:  []string{},
			output: "",
		},
		{
			name:   "empty string",
			input:  []string{"", "cat", "cattle"},
			output: "",
		},
		{
			name:   "nil input",
			input:  nil,
			output: "",
		},
		{
			name:   "single stirng",
			input:  []string{"single"},
			output: "single",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := longestCommonPrefix(tc.input)
			if tc.output != output {
				t.Errorf("input: %v, ouput: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	r := strs[0]
rangestrs:
	for _, str := range strs[1:] {
		for r != "" {
			if len(str) < len(r) {
				r = r[:len(str)]
			}
			if str[:len(r)] == r {
				continue rangestrs
			}
			r = r[:len(r)-1]
		}
		return ""
	}
	return r
}

// @lc code=end
