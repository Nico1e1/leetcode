package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) set(val int) {
	node.Val = val
}

func (node *TreeNode) get() int {
	return node.Val
}

func (node TreeNode) setVal(val int) {
	node.Val = val
}
