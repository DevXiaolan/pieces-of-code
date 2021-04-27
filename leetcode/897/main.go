package main

import "fmt"
type TreeNode struct {
	Val int
	Left  *TreeNode
	Right *TreeNode
}
var Left = -1
var Right = 1



func rFind(root *TreeNode,sig int) *TreeNode {
	if sig == 1 {
		//right
		if root.Right == nil {
			return root
		}
		return rFind(root.Right, sig)
	}else {
		//left
		if root.Left == nil {
			return root
		}
		return rFind(root.Left, sig)
	}
}


func increasingBST(root *TreeNode) *TreeNode {
	l:=root
	if root.Right!=nil {
		root.Right =increasingBST(root.Right)
	}
	if root.Left != nil {
		l = increasingBST(root.Left)
		rFind(l,Right).Right = root
	}
	return l
}

func pp(root *TreeNode) {
	for root != nil {
		fmt.Printf("%+v\n",[]interface{}{root.Val})
		root = root.Right
	}
}

func main(){
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:3,
			Left: &TreeNode{
				Val:2,
				Left: &TreeNode{Val:1},
			},
			Right: &TreeNode{
				Val:4,		
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{Val:7},
				Right: &TreeNode{Val:9},
			},
		},
	}
	pp(increasingBST(root))
}