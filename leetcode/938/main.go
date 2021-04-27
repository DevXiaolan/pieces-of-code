package main

import "fmt"


type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if root == nil {
		return sum
	}
	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	return sum + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
func main(){
	//[10,5,15,3,7,null,18]
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:5,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:15,
			Right: &TreeNode{Val:18},
		},
	}
	fmt.Printf("%+v\n",[]interface{}{rangeSumBST(root,7,15)})
}