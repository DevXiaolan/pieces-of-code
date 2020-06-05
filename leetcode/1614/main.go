package main

import (
	"fmt"
)

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 1 {
			i++
		} else if nums[j]%2 == 0 {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	return nums
}

func main() {
	fmt.Printf("%+v\n", []interface{}{exchange([]int{1, 2, 3, 4})})
}
