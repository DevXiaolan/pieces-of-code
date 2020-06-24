package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func Print(root *ListNode) {
	if root != nil {
		fmt.Printf("%+v\n", []interface{}{root.Val})
		Print(root.Next)
	}
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := head
	for i := 0; i < k; i++ {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		slow = slow.Next
	}
	return slow
}
func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	Print(getKthFromEnd(root, 2))
}
