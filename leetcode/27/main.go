package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	f := removeElement(nums, 3)
	fmt.Printf("%+v\n", []interface{}{f, nums})
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}
	found := 0
	for k := 0; k < len(nums); k++ {
		v := nums[k]
		if v == val {
			found++
			continue
		}
		nums[k-found] = v
	}
	return len(nums) - found
}
