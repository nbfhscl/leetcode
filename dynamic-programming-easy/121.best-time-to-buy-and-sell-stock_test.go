package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMaxProfit(t *testing.T) {
	myt := easytest.New(t).F(maxProfit)
	myt.Run([]int{7, 1, 5, 3, 6, 4}).Equal(5)
	myt.Run([]int{7, 6, 4, 3, 1}).Equal(0)
	myt.Run([]int{7}).Equal(0)
	myt.Run([]int{0, 0}).Equal(0)
	myt.Run([]int{0, 1}).Equal(1)
}

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start

// profit[i] = max profit you can get on day i
// if profit[i-1] < 0, profit[i] = prices[i] - profit[i-1]
// else profit[i] += prices[i] - profit[i-1]
func maxProfit(prices []int) int {
	var max int = 0
	var last int = prices[0]
	var lastProfit int = 0
	for i := 1; i < len(prices); i++ {
		if lastProfit > 0 {
			lastProfit += prices[i] - last
		} else {
			lastProfit = prices[i] - last
		}
		if lastProfit > max {
			max = lastProfit
		}
		last = prices[i]
	}
	return max
}

// @lc code=end
