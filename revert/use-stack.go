package main

import "lanhao/revert/util"

func RevertWithStack(current util.Node) util.Node {
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
