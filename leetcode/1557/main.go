package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	max := -1 * 2 << 31
	tmp := nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	for _, v := range nums {
		if tmp+v < 0 {
			if tmp > max {
				max = tmp
			}
			tmp = v
		} else {
			tmp += v
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxSubArray([]int{-2, -1})})
}
