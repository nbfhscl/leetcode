package test

import (
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestClimbStairs(t *testing.T) {
	myt := easytest.New(t)
	myt.F(climbStairs2).Run(1).Equal(1)
	myt.Run(2).Equal(2)
	myt.Run(3).Equal(3)
	myt.Run(4).Equal(5)
	myt.Run(45).Print()
}

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
// @lc code=start

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	vi, vj := 1, 2
	k, vk := 3, 0
	for k <= n {
		vk = vi + vj
		vi = vj
		vj = vk
		k++
	}
	return vk
}

// @lc code=end

func climbStairs2(n int) int {
	m := make(map[int]int)
	return climb(n, m)
}

func climb(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if _, ok := m[n]; !ok {
		m[n] = climb(n-1, m) + climb(n-2, m)
	}
	return m[n]
}
