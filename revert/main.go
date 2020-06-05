package main

import (
	"fmt"
	"lanhao/revert/util"
)

func main() {
	node := util.InitNode(4)
	util.Print(node)
	// nOut := revertWithStack(node)
	nOut := util.Node{}
	revertWithRecursion(node, &nOut)
	fmt.Printf("%+v", nOut)
	util.Print(nOut)
}

/**
* 使用递归实现
 */
func revertWithRecursion(current util.Node, prev util.Node, p *util.Node) {

	if current.Next == nil {
		p = current.Next
		p.Next = &current
	}

}

/**
* 使用栈实现
 */
func revertWithStack(current util.Node) util.Node {
	stack := []util.Node{}
	pointer := &current
	for {
		stack = append(stack, *pointer)
		if pointer.Next == nil {
			break
		}
		pointer = pointer.Next
	}

	nOut := stack[len(stack)-1]
	pointer = &nOut
	for i := len(stack) - 2; i >= 0; i-- {
		pointer.Next = &stack[i]
		pointer = pointer.Next
		if i == 0 {
			pointer.Next = nil
		}
	}
	return nOut
}
