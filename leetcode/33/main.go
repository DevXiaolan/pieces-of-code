package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	n := search(nums, 3)
	fmt.Printf("%+v\n", []interface{}{n})
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] > target {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	}
	return -1
}
