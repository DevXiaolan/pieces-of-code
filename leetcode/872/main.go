package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaf(root *TreeNode) []int {
	leaf := []int{}
	if root != nil {
		if root.Left == nil && root.Right == nil {
			leaf = append(leaf, root.Val)
		}

		leaf = append(leaf, findLeaf(root.Left)...)
		leaf = append(leaf, findLeaf(root.Right)...)
	}
	return leaf
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	lf1 := findLeaf(root1)
	fmt.Printf("%+v\n", []interface{}{lf1})
	lf2 := findLeaf(root2)
	fmt.Printf("%+v\n", []interface{}{lf2})
	if len(lf1) != len(lf2) {
		return false
	}
	for k := 0; k < len(lf1); k++ {
		if lf1[k] != lf2[k] {
			return false
		}
	}
	return true
}

func main() {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	// [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Printf("%+v\n", []interface{}{leafSimilar(root1, root2)})
}
