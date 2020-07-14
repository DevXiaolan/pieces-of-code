package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func getDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getDeep(root.Left), getDeep(root.Right))
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	diff := getDeep(root.Left) - getDeep(root.Right)

	return diff*diff <= 1
}
func main() {
	root := &TreeNode{
		Val: 1,
	}
	// [1,2,2,3,null,null,3,4,null,null,4]
	// root := &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 9,
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 20,
	// 		Left: &TreeNode{
	// 			Val: 15,
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 			Left: &TreeNode{
	// 				Val: 1,
	// 			},
	// 		},
	// 	},
	// }
	fmt.Printf("%+v\n", []interface{}{isBalanced(root)})
}
