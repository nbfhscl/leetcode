package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestThreeSum(t *testing.T) {
	myt := easytest.New(t).F(threeSum)
	myt.Input([]int{}).Output([][]int{})
	myt.Input([]int{0}).Output([][]int{})
	myt.Input([]int{0, 0}).Output([][]int{})
	myt.Input([]int{0, 0, 0}).Output([][]int{{0, 0, 0}})
	myt.Input([]int{0, 0, 0, 0}).Output([][]int{{0, 0, 0}})
	myt.Input([]int{0, 0, 0, 1, -1, 1, -1, 0}).Output([][]int{{-1, 0, 1}, {0, 0, 0}})
	myt.Input([]int{-1, 0, 1, 2, -1, -4}).Output([][]int{{-1, -1, 2}, {-1, 0, 1}})
}

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

// 28ms
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == -nums[i] {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l, r = l+1, r-1
			} else if nums[l]+nums[r] < -nums[i] {
				l++
			} else if nums[l]+nums[r] > -nums[i] {
				r--
			}
		}

	}
	return res
}

// @lc code=end

// 452ms
func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			v := nums[j]
			if _, ok := m[v]; !ok {
				m[-nums[i]-v] = v
				continue
			}
			res = append(res, []int{nums[i], m[v], v})
			for j < len(nums)-1 && nums[j] == nums[j+1] {
				j++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
