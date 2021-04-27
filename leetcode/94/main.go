package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var r []int

func zx(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		zx(root.Left)
	}
	r = append(r, root.Val)
	if root.Right != nil {
		zx(root.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	r = []int{}
	zx(root)
	return r
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Printf("%+v\n", []interface{}{inorderTraversal(root)})
}
