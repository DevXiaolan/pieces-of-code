package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	min, max := nums[0], nums[0]
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		sum += nums[i]
	}
	miss := (min+max)*(max-min+1)/2 - sum
	if miss == 0 {
		miss = max+1
	}
	return miss
}
func main() {
	fmt.Printf("%+v\n", []interface{}{missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})})
}
