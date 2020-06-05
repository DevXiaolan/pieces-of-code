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

var lays []float64
var count []int

func dig(root *TreeNode, layer int) {
	if root == nil {
		return
	}
	if len(lays) <= layer {
		lays = append(lays, 0)
		count = append(count, 0)
	}
	lays[layer] = (lays[layer]*float64(count[layer]) + float64(root.Val)) / float64(count[layer]+1)
	count[layer]++
	// fmt.Printf("%+v\n",[]interface{}{lays,count})
	dig(root.Left, layer+1)
	dig(root.Right, layer+1)
}

func averageOfLevels(root *TreeNode) []float64 {
	lays = []float64{}
	count = []int{}
	dig(root, 0)
	return lays
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Printf("%+v\n", []interface{}{averageOfLevels(root)})
}
