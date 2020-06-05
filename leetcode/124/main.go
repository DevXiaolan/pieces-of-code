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

var sum int

/**
* 思路
* 从顶点往下找，当路径经过这个顶点时，路径和为: 左 + 中 + 右
* 其中，左、右如果小于0，实际上相当于不走这个路径，所以路径和计算0就可以
* 然后递归下去，如果子节点的路径和大于父节点，那这个路径就不会经过父节点，
* 以上
 */
func r(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 经过当前顶点的路径和 中+左+右
	left := r(root.Left)
	if left < 0 {
		left = 0
	}
	right := r(root.Right)
	if right < 0 {
		right = 0
	}
	nodeSum := root.Val + left + right

	if nodeSum > sum {
		sum = nodeSum
	}
	// 返回这个节点  如果作为路径的“一部分”的和
	// 如果作为一部分，路径不可能同时经过当前节点的左和右
	// 所以，只返回 左、右 的较大一个
	tailVal := right
	if left > right {
		tailVal = left
	}
	return root.Val + tailVal
}

func maxPathSum(root *TreeNode) int {
	sum = -999999999999
	r(root)
	return sum
}

func main() {
	
	node := TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: 2,
		},
		// Right:&TreeNode{
		// 	Val:3,
		// 	// Left:&TreeNode{
		// 	// 	Val:15,
		// 	// },
		// 	// Right:&TreeNode{
		// 	// 	Val:7,
		// 	// },
		// },
	}
	
	fmt.Printf("%+v\n", []interface{}{maxPathSum(&node)})
}
