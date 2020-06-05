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
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	out := swapPairs(n)
	for out != nil {
		fmt.Printf("%+v\n", []interface{}{out.Val})
		out = out.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	p := head
	swap := false
	stack := []*ListNode{}
	for p != nil {
		stack = append(stack, p)
		p = p.Next
		swap = !swap
	}

	out := ListNode{}
	pOut := &out
	if len(stack) == 1 {
		pOut.Next = stack[0]
		return out.Next
	}
	for i := 0; i < len(stack); i++ {
		if i%2 == 0 {
			if i+1 < len(stack) {
				continue
			} else {
				pOut.Next = stack[i]
			}
		} else {
			pOut.Next = stack[i]
			pOut.Next.Next = stack[i-1]
			pOut = pOut.Next.Next
			pOut.Next = nil
		}
	}
	return out.Next
}
