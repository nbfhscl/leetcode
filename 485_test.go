/*
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
	Input: [1,1,0,1,1,1]
	Output: 3
	Explanation: The first two digits or the last three digits are consecutive 1s.
		The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.

The length of input array is a positive integer and will not exceed 10,000
*/

package test

import (
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tc := []struct {
		name string
		nums []int
		output int
	}{
		{"example", []int{1,1,0,1,1,1}, 3},
		{"first", []int{1,1,1,1,1,1,1,0,0,1,1,0,1,1,0,1,1,1,1,1,1,0}, 7},
	}
	for i := range tc {
		t.Run(tc[i].name, func(t *testing.T){
			result := findMaxConsecutiveOnes(tc[i].nums)
			if result != tc[i].output {
				t.Fatal(result)
			}
		})
	}
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	curr := 0
	for i :=range nums {
		if nums[i] == 1 {
			curr ++
		} else {
			if curr > max {
				max = curr
			}
			curr = 0
		}
	}
	if curr > max {
		max = curr
	}
	return max
}