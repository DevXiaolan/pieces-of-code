package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	l := len(nums)
	if l == 0 {
		return []int{0}
	}
	if l == 1 {
		return []int{nums[0]}
	}
	if l == 2 {
		return []int{nums[0] + nums[1]}
	}
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 {
			result = append(result, nums[0]+nums[1]+nums[2])
		} else {
			result = append(result, result[i-1]-nums[i-1]+nums[i+2])
		}
	}
	return result
}

func main() {
	fmt.Printf("%+v\n", []interface{}{maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)})
}
