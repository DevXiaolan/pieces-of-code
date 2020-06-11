package main

import (
	"fmt"
)

// Stack -
type Stack struct {
	queue  []int
	length int
}

func (s *Stack) push(i int) {
	s.queue = append(s.queue, i)
	s.length++
}
func (s *Stack) pop() {
	s.queue = s.queue[:len(s.queue)-1]
	s.length--
}
func (s Stack) print() {
	fmt.Printf("%+v\n", []interface{}{s.queue})
}
func (s Stack) empty() bool {
	return s.length == 0
}
func (s Stack) last() int {
	return s.queue[s.length-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := &Stack{}
	l := len(popped)
	i :=0
	for _, el := range pushed {
		stack.push(el)
		for i < l && !stack.empty() && stack.last() == popped[i] {
			stack.pop()
			i++
		}
	}
	return stack.empty()
}

func main() {
	fmt.Printf("%+v\n", []interface{}{validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})})
}
