package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestInsert(t *testing.T) {
	myt := easytest.New(t).F(insert)
	myt.Input([][]int{{1, 3}, {2, 5}}, []int{2, 5}).Output([][]int{{1, 5}})
	myt.Input([][]int{}, []int{5, 7}).Output([][]int{{5, 7}})
	myt.Input([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}).Output([][]int{{1, 2}, {3, 10}, {12, 16}})

	myt.Input([][]int{{5, 6}, {7, 8}}, []int{1, 2}).Output([][]int{{1, 2}, {5, 6}, {7, 8}})
	myt.Input([][]int{{5, 6}, {7, 8}}, []int{9, 10}).Output([][]int{{5, 6}, {7, 8}, {9, 10}})
	myt.Input([][]int{{5, 6}, {7, 8}}, []int{5, 8}).Output([][]int{{5, 8}})
}

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start

//              [newInterval]
//   [left] -> right < newInterval.left
//                              [right] -> left > newInterval.right
// other intervals is in middle and should overlapping with newInterval
func insert(intervals [][]int, newInterval []int) [][]int {
	var i int
	res := make([][]int, 0, len(intervals))
	for i = 0; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		if newInterval[0] > intervals[i][0] {
			newInterval[0] = intervals[i][0]
		}
		if newInterval[1] < intervals[i][1] {
			newInterval[1] = intervals[i][1]
		}
	}
	res = append(res, newInterval)
	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}

// @lc code=end
