package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Printf("%+v\n", []interface{}{nums})
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := nums[0]
	length := 1
	for _, v := range nums {
		if last == v {
			continue
		}
		nums[length] = v
		last = v
		length++
	}
	return length
}
