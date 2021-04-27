package main

import "fmt"


 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func convOne(root *TreeNode, prev int) int {
	sum := 0
	if root == nil {
		return 0
	}
	if root.Right == nil {
		root.Val += prev
	}
	if root.Right != nil {
		rightSum := convOne(root.Right, prev)
		root.Val += rightSum
		sum += rightSum
	}
	if root.Left != nil {
		leftSum := convOne(root.Left,root.Val)
		sum += leftSum
	}
	return sum + root.Val
}

func convertBST(root *TreeNode) *TreeNode {
	convOne(root,0)
	return root
}

func main(){
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:-1,
			Left: &TreeNode{
				Val:-3,
			},
		},
		Right: &TreeNode{
			Val:2,
			Right: &TreeNode{
				Val:4,
			},
		},
	}
	r := convertBST(root)
	fmt.Printf("%+v\n",[]interface{}{
		r.Val,
		r.Left.Val,
		r.Left.Left.Val,
	})
}