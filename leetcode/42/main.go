package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", []interface{}{trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})})
}

type Stack struct {
	P int
	D []int
}

func trap(height []int) int {
	result := 0
	hLen := len(height)
	stack := Stack{
		P: 0,
		D: []int{0},
	}
	for i := 1; i < hLen; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			level := height[stack[len(stack)-1]]
			stack = pop(stack)
			if len(stack) == 0 {
				break
			}
			result += (i - stack[len(stack)-1] - 1) * (min(height[i], height[stack[len(stack)-1]]) - level)
		}
		stack = append(stack, i)
	}
	return result
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
