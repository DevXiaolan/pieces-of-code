package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
			// &ListNode{
			// 	Val: 3,
			// 	Next: &ListNode{
			// 		Val: 4,
			// 		Next: &ListNode{
			// 			Val:  5,
			// 			Next: nil,
			// 		},
			// 	},
			// },
		},
	}
	out := removeNthFromEnd(n, 2)
	for out != nil {
		fmt.Printf("%+v\n", []interface{}{out.Val})
		out = out.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	step := 1
	pAhead := head
	var pTargetPrev *ListNode
	var pTarget *ListNode

	for pAhead != nil {
		if step == n {
			pTarget = head
		}
		if step == n+1 {
			pTargetPrev = head
			pTarget = pTargetPrev.Next
		}
		if step > n+1 {
			pTargetPrev = pTargetPrev.Next
			pTarget = pTargetPrev.Next
			fmt.Printf("%+v\n", []interface{}{step, *pTargetPrev, *pTarget})
		}
		step++
		pAhead = pAhead.Next
	}

	if pTargetPrev != nil {
		if pTarget != nil {
			pTargetPrev.Next = pTarget.Next
		} else {
			pTargetPrev.Next = nil
		}
	} else if pTargetPrev == nil && pTarget != nil {
		head = pTarget.Next
	} else {
		head = nil
	}

	return head
}
