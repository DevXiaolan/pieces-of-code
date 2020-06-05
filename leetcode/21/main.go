package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Printf("%+v\n", []interface{}{l3.Val})
		l3 = l3.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	if p1 == nil && p2 == nil {
		return nil
	}
	var l3 = ListNode{
		Val:  0,
		Next: nil,
	}
	p3 := &l3
	for {
		if p1 == nil && p2 == nil {
			break
		}
		if p1 == nil {
			p3.Val = p2.Val
			p2 = p2.Next
			if p2 != nil {
				p3.Next = &ListNode{
					Val:  0,
					Next: nil,
				}
				p3 = p3.Next
			}
		} else if p2 == nil {
			p3.Val = p1.Val
			p1 = p1.Next
			if p1 != nil {
				p3.Next = &ListNode{
					Val:  0,
					Next: nil,
				}
				p3 = p3.Next
			}
		} else if p1.Val > p2.Val {
			p3.Val = p2.Val
			p3.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			p2 = p2.Next
			p3 = p3.Next
		} else if p1.Val <= p2.Val {
			p3.Val = p1.Val
			p3.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			p1 = p1.Next
			p3 = p3.Next
		}
	}

	return &l3
}
