package main

import (
	"fmt"
)

// TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func r(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return nil
	}
	if d > 2 {
		// 递归
		root.Left = r(root.Left, v, d-1)
		root.Right = r(root.Right, v, d-1)
		return root
	} else if d == 2 {
		// 寄存起来
		tLeft := root.Left
		tRight := root.Right

		newLeftNode := TreeNode{
			Val:  v,
			Left: tLeft,
		}
		root.Left = &newLeftNode

		newRightNode := TreeNode{
			Val:   v,
			Right: tRight,
		}
		root.Right = &newRightNode
		return root
	} else {

		newNode := TreeNode{
			Val:  v,
			Left: root,
		}
		root = &newNode
		return root
	}
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	newRoot := r(root, v, d)
	return newRoot
}

func main() {
	// 4,2,6,3,1,5
	node := TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	addOneRow(&node, 1, 1).toString()
}

func (node TreeNode) toString() {
	dig(&node, true)
}

func dig(node *TreeNode, top bool) {
	if node != nil {
		if top {
			fmt.Println(node.Val)
		}
		if node.Left != nil {
			fmt.Println(node.Left.Val)
		}
		if node.Right != nil {
			fmt.Println(node.Right.Val)
		}
		dig(node.Left, false)
		dig(node.Right, false)
	}
}
