package main

import (
	"fmt"
)

// Node -
type Node struct {
	Val      int
	Children []*Node
}

var result []int

func cl(root *Node) {
	if root == nil {
		return
	}
	if root.Children != nil {
		for _, c := range root.Children {
			cl(c)
		}
	}
	result = append(result, root.Val)
}

func postorder(root *Node) []int {
	result = []int{}
	cl(root)
	return result
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}
	fmt.Printf("%+v\n",[]interface{}{postorder(root)})
}
