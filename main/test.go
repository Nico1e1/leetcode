package main

import "fmt"

func main() {
	node := &TreeNode{Val: 0}
	node.setVal(1)
	fmt.Println(node.get())
}
