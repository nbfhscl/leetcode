package test

import (
	"math"
	"testing"

	"git.letmeout.cn/nbfhscl/codes/go/easytest"
)

func TestMinStack(t *testing.T) {
	myt := easytest.New(t)
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	myt.Equal(minstack.GetMin(), -3)
	minstack.Pop()
	myt.Equal(minstack.Top(), 0)
	myt.Equal(minstack.GetMin(), -2)
}

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	nums []int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nums: make([]int, 0, 16), min: math.MaxInt64}
}

// if x > min, append x, 1,2,3,4,5
// if x <= min, append min and x
func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.nums = append(this.nums, this.min)
		this.min = x
	}
	this.nums = append(this.nums, x)
}

// if top != min, remove top
// if top == min, remove top, set min to top, remove top
func (this *MinStack) Pop() {
	top := this.nums[len(this.nums)-1]
	if top == this.min {
		this.nums = this.nums[:len(this.nums)-1]
		this.min = this.nums[len(this.nums)-1]
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
