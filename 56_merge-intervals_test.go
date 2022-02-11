package test

import (
	"sort"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMerge(t *testing.T) {
	myt := easytest.New(t).F(merge)
	myt.Input([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}).Output([][]int{{1, 6}, {8, 10}, {15, 18}})
	myt.Input([][]int{{4, 5}, {1, 4}}).Output([][]int{{1, 5}})
	myt.Input([][]int{}).Output([][]int{})
	myt.Input([][]int{{0, 1}}).Output([][]int{{0, 1}})
	myt.Input([][]int{{0, 0}, {0, 0}}).Output([][]int{{0, 0}})
	myt.Input([][]int{{0, 0}, {0, 1}}).Output([][]int{{0, 1}})
	myt.Input([][]int{{0, 1}}).Output([][]int{{0, 1}})

	// special cases
	myt.Input([][]int{{1, 4}, {2, 3}}).Output([][]int{{1, 4}})
}

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start

type Interval [][]int

func (itv Interval) Swap(i, j int)      { itv[i], itv[j] = itv[j], itv[i] }
func (itv Interval) Less(i, j int) bool { return itv[i][0] < itv[j][0] }
func (itv Interval) Len() int           { return len(itv) }

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(Interval(intervals))
	intervals = append(intervals, nil)
	cnt := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i] != nil && intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
			if intervals[i][1] < intervals[i-1][1] {
				intervals[i][1] = intervals[i-1][1]
			}
		} else {
			intervals[cnt] = intervals[i-1]
			cnt++
		}
	}
	return intervals[:cnt]
}

// @lc code=end
