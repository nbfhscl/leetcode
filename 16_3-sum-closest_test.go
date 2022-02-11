package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestThreeSumClosest(t *testing.T) {
	myt := easytest.New(t).F(threeSumClosest)
	myt.Input([]int{-1, 2, 1, -4}, 1).Output(2)
	myt.Input([]int{0, 0, 0, 1}, 100).Output(1)
}

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	clo := 10001
	neg := false
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			diff := nums[l] + nums[r] + nums[i] - target
			if diff < 0 {
				if -diff < clo {
					clo = -diff
					neg = true
				}
				l++
			} else if diff > 0 {
				if diff < clo {
					clo = diff
					neg = false
				}
				r--
			} else if diff == 0 {
				return target
			}
		}
	}
	if neg {
		clo = target - clo
	} else {
		clo = target + clo
	}
	return clo
}

// @lc code=end
