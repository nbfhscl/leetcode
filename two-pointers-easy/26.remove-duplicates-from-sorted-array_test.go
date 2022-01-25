package test

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "no duplicated",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "all same",
			input:  []int{1, 1, 1, 1, 1},
			output: []int{1},
		},
		{
			name:   "example1",
			input:  []int{1, 1, 2},
			output: []int{1, 2},
		},
		{
			name:   "example2",
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: []int{0, 1, 2, 3, 4},
		},
		{
			name:   "empty",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "nil",
			input:  nil,
			output: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.input[:removeDuplicates(tc.input)]
			if !reflect.DeepEqual(tc.output, output) {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	var r int = 1
	for _, v := range nums[1:] {
		if v == nums[r-1] {
			continue
		}
		nums[r] = v
		r++
	}
	return r
}

// @lc code=end

func removeDuplicates2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	var r int
	for _, v := range nums {
		if m[v] {
			continue
		}
		nums[r] = v
		r++
		m[v] = true
	}
	return r
}
