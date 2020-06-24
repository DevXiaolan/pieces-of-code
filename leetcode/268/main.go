package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	sum, max, min := 0, 0, 2<<32

	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	if min == 1 {
		return 0
	}
	diff := (max+min)*(max-min+1)/2 - sum
	if diff ==0 {
		// 完整数列
		return max+1
	}
	return diff
}

func main() {
	fmt.Printf("%+v\n", []interface{}{missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})})
}
