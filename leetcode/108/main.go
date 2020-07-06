package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	i := len(nums) / 2
	mid := nums[i]
	// fmt.Printf("%+v\n", []interface{}{mid, nums[:i], nums[i+1:]})
	return &TreeNode{
		Val:   mid,
		Left:  sortedArrayToBST(nums[:i]),
		Right: sortedArrayToBST(nums[i+1:]),
	}
}

func main() {
	fmt.Printf("%+v\n", []interface{}{sortedArrayToBST([]int{-10, -3, 0, 5, 9})})
}
