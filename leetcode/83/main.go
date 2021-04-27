package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := map[int]bool{
		head.Val: true,
	}
	prev := head
	p := head
	for p != nil {
		_, ok := m[p.Val]
		for ok {
			prev.Next = p.Next
			p = p.Next
			if p == nil {
				return head
			}
			_, ok = m[p.Val]
		}
		m[p.Val] = true
		prev = p
		p = p.Next
	}

	return head
}
func main() {
	// link := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 1,
	// 		Next: &ListNode{
	// 			Val: 2,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 3,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	r := deleteDuplicates(nil)
	for r != nil {
		fmt.Printf("%+v\n", []interface{}{r.Val})
		r = r.Next
	}
}
