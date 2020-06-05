package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var strs []string

func p(prefix string, root *TreeNode) {
	if root == nil {
		return
	}
	var str string
	if len(prefix) > 0 {
		str = prefix + "->" + strconv.Itoa(root.Val)
	} else {
		str = strconv.Itoa(root.Val)
	}

	if root.Left != nil {
		p(str, root.Left)
	}
	if root.Right != nil {
		p(str, root.Right)
	}
	if root.Left == nil && root.Right == nil {
		strs = append(strs, str)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	strs = []string{}
	if root == nil {
		return strs
	}
	p("", root)
	return strs
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Printf("%+v\n", []interface{}{binaryTreePaths(root)})
}
