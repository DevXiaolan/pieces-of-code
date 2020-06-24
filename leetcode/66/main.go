package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for {
		if digits[i] == 9 {
			digits[i] = 0
			i--
			if i < 0 {
				return append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}

func main() {
	fmt.Printf("%+v\n", []interface{}{plusOne([]int{9, 9, 9})})
}
