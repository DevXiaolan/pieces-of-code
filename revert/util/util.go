package util

import "fmt"

type Node struct {
	V    uint32
	Next *Node
}

func InitNode(n int) Node {
	node := Node{
		V: 1,
	}

	pointer := &node
	for i := 0; i < n-1; i++ {
		next := Node{}
		next.V = pointer.V + 1
		pointer.Next = &next
		pointer = &next
	}
	return node
}

func Print(node Node) {
	pointer := &node
	for {
		fmt.Println(pointer.V)
		if pointer.Next == nil {
			break
		}
		pointer = pointer.Next
	}

}
