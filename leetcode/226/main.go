package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func main() {
	root := &TreeNode{
		Val:   4,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	fmt.Printf("%+v\n", []interface{}{invertTree(root)})
}
