package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for l != nil {
		fmt.Printf("%+v\n", []interface{}{l.Val})
		l = l.Next
	}
}

func revert(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p *ListNode = nil
	g := head
	for g != nil {
		tmp := g
		g = g.Next
		if p == nil {
			tmp.Next = nil
		}
		tmp.Next = p
		p = tmp
	}
	return p
}

// 快慢指针？
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for fast != nil {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
		if fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
		slow = slow.Next
	}
	tailPart := revert(slow.Next)
	for tailPart != nil {
		
		if tailPart.Val != head.Val {
			return false
		}
		tailPart = tailPart.Next
		head = head.Next
	}
	return true
}

// var h *ListNode

// func foo(node *ListNode) bool {
// 	sig := true

// 	if node.Next != nil {
// 		sig = foo(node.Next)
// 	}

// 	if sig && node.Val == h.Val {
// 		h = h.Next
// 		return true
// 	}
// 	return false
// }
// // 递归法
// func isPalindrome(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return true
// 	}
// 	h = head
// 	return foo(head)
// }

// 数组法
// func isPalindrome(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return true
// 	}
// 	arr := []int{}
// 	for head != nil {
// 		arr = append(arr, head.Val)
// 		head = head.Next
// 	}
// 	i, j := 0, len(arr)-1
// 	for i < j {
// 		if arr[i]!=arr[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Printf("%+v\n", []interface{}{isPalindrome(root)})
}
