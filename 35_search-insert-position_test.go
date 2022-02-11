package test

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		target int
		output int
	}{
		{
			name:   "example 1",
			input:  []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			name:   "example 1",
			input:  []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			name:   "example 3",
			input:  []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
		{
			name:   "example 4",
			input:  []int{1, 3, 5, 6},
			target: 0,
			output: 0,
		},
		{
			name:   "empty input",
			input:  []int{},
			output: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := searchInsert(tc.input, tc.target)
			if tc.output != output {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
