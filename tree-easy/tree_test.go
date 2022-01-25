package test

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(ns ...interface{}) *TreeNode {
	if len(ns) <= 1 {
		return nil
	}
	return newTree(ns, 1)
}

func newTree(ns []interface{}, index int) *TreeNode {
	var tn *TreeNode
	if len(ns) > index && ns[index] != nil {
		tn = &TreeNode{Val: ns[index]}
	} else {
		return nil
	}
	tn.Left = newTree(ns, 2*index)
	tn.Right = newTree(ns, 2*index+1)
	return tn
}

func newTree1(ns []int, index int) *TreeNode {
	var tn *TreeNode
	if len(ns) > index {
		tn = &TreeNode{Val: ns[index]}
	} else {
		return nil
	}
	tn.Left = newTree1(ns, 2*index+1)
	tn.Right = newTree1(ns, 2*index+2)
	return tn
}
