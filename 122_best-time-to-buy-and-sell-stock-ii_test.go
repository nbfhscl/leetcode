package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMaxProfit(t *testing.T) {
	myt := easytest.New(t).F(maxProfit)
	myt.Run([]int{7, 1, 5, 3, 6, 4}).Equal(7)
	myt.Run([]int{1, 2, 3, 4, 5}).Equal(4)
	myt.Run([]int{7, 6, 4, 3, 1}).Equal(0)
	myt.Run([]int{1}).Equal(0)
}

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start

// profit[i][0] = max profit you can get on day i with 0 stock on hand
// profit[i][1] = max profit you can get on day i with 1 stock on hand
// profit[i][0] = max( profit[i-1][0], profit[i-1][1] + prices[i] )
// profit[i][1] = max( profit[i-1][1], profit[i-1][0] - prices[i] )
func maxProfit(prices []int) int {
	profit0 := 0
	profit1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		temp := max(profit0, profit1+prices[i])
		profit1 = max(profit1, profit0-prices[i])
		profit0 = temp
	}
	return profit0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end
