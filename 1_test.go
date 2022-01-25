/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

    Input: nums = [2,7,11,15], target = 9
    Output: [0,1]
    Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

    Input: nums = [3,2,4], target = 6
    Output: [1,2]
Example 3:

    Input: nums = [3,3], target = 6
    Output: [0,1]
*/
package test

import (
    "testing"
)

func TestTwoSum(t *testing.T) {
    tc := []struct{
        name string
        nums []int
        target int
    }{
        {"one", []int{2,7,11,15},9},
        {"two", []int{3,2,4},6},
        {"three", []int{3,3},6},
    }
    for i := range tc {
        t.Run(tc[i].name, func(t *testing.T) {
            if result := twoSum2(tc[i].nums, tc[i].target); result == nil || (tc[i].nums[result[0]] + tc[i].nums[result[1]] != tc[i].target) {
                t.Fatal()
            }
        })
    }
}

func twoSum(nums []int, target int) []int {
    length := len(nums)
    for i:=0;i<length;i++ {
        for j:=i+1;j<length;j++ {
            if nums[i] + nums[j] == target {
                return []int{i,j}
            }
        }
    }
    return nil
}

func twoSum2(nums []int, target int) []int {
    length := len(nums)
    m := map[int]int{}
    for i:=0;i<length;i++ {
        cmpt := target - nums[i]
        if _, ok := m[cmpt]; ok {
            return []int{m[cmpt], i}
        }
        m[nums[i]] = i
    }
    return nil
}