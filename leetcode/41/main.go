package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", []interface{}{firstMissingPositive([]int{3, 1})})
}

func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 1
	}
	for i := 0; i < numsLen; i++ {
		// fmt.Printf("%+v\n", []interface{}{i, nums[i] > 0 && nums[i] != i+1 && nums[i] < numsLen})
		for nums[i] > 0 && nums[i] != i+1 && nums[i] < numsLen {
			swap(nums, nums[i]-1, i)
			if nums[i] > 0 && nums[i] < numsLen && nums[i] == nums[nums[i]-1] {
				break
			}
		}
	}
	for i := 0; i < numsLen; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return numsLen + 1
}

func swap(nums []int, i int, j int) {
	nums[i] = nums[i] ^ nums[j]
	nums[j] = nums[i] ^ nums[j]
	nums[i] = nums[i] ^ nums[j]
}
