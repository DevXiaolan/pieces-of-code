package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	stack := []int{}
	for i := 0; i < k; i++ {
		stack = append(stack, -1*(2<<31))
	}
	for _, v := range nums {
		for i := 0; i < k; i++ {
			if v > stack[i] {
				v = v ^ stack[i]
				stack[i] = v ^ stack[i]
				v = v ^ stack[i]
			}
		}
	}
	fmt.Printf("%+v\n", []interface{}{stack})
	return stack[k-1]
}

func main() {
	fmt.Printf("%+v\n", []interface{}{findKthLargest([]int{-1, -1}, 2)})
}
