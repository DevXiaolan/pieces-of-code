package main

import "fmt"

// TreeNode d
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricArr(arr []*TreeNode) bool {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] == nil || arr[l-i-1] == nil {
			if arr[i] != arr[l-i-1] {
				return false
			}
		} else {
			if arr[i].Val != arr[l-i-1].Val {
				return false
			}
		}
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	layers := []*TreeNode{
		root,
	}
	sig := true
	for len(layers) > 0 {
		sig = sig && isSymmetricArr(layers)
		newL := []*TreeNode{}
		for _, item := range layers {
			if item != nil {
				newL = append(newL, item.Left)
				newL = append(newL, item.Right)
			} else {
				newL = append(newL, nil, nil)
			}
		}
		layers = newL
		if !sig {
			return sig
		}
	}
	return true
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	fmt.Printf("%+v\n", []interface{}{isSymmetric(root)})
}
