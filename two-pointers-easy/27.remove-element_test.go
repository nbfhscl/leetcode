package test

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tcs := []struct {
		name   string
		input  []int
		input2 int
		output []int
	}{
		{
			name:   "return all",
			input:  []int{1, 2, 3, 4, 5},
			input2: 0,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "return empty",
			input:  []int{1, 1, 1, 1},
			input2: 1,
			output: []int{},
		},
		{
			name:   "example1",
			input:  []int{3, 2, 2, 3},
			input2: 3,
			output: []int{2, 2},
		},
		{
			name:   "example2",
			input:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			input2: 2,
			output: []int{0, 1, 4, 0, 3},
		},
		{
			name:   "empty",
			input:  []int{},
			input2: 100,
			output: []int{},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.input[:removeElement(tc.input, tc.input2)]
			if !reflect.DeepEqual(tc.output, output) {
				t.Errorf("input: %v, output: %v, expected: %v", tc.input, output, tc.output)
			}
		})
	}
}

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 || nums == nil {
		return 0
	}
	var head, tail int
	tail = len(nums) - 1
	for head < tail {
		for head < tail {
			if nums[head] == val {
				break
			}
			head++
		}
		for head < tail {
			if nums[tail] != val {
				nums[head] = nums[tail]
				tail--
				break
			}
			tail--
		}
	}
	if nums[head] == val {
		return head
	}
	return head + 1
}

// @lc code=end
