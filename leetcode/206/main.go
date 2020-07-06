package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var ll ListNode
var p *ListNode

func doRevert(head *ListNode) {
	if head == nil {
		return
	}
	if head.Next != nil {
		doRevert(head.Next)
	}
	p.Next = head
	head.Next = nil
	p = head
}
func reverseList(head *ListNode) *ListNode {
	ll = ListNode{}
	p = &ll
	doRevert(head)
	return ll.Next
}

func Print(root *ListNode) {
	for root != nil {
		fmt.Printf("%+v\n", []interface{}{root.Val})
		root = root.Next
	}
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
	Print(reverseList(root))
}
