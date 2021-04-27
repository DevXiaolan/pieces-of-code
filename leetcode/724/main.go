package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, 0
	for _, v := range nums {
		right += v
	}

	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if i > 0 {
			left += nums[i-1]
		}
		if left == right {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Printf("%+v\n", []interface{}{pivotIndex([]int{-1, -1, -1, 0, 1, 1})})
}
