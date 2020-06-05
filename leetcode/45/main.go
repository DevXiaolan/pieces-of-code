package main

import "fmt"

func main() {
	// a, b := rangeMax([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1}, 1)
	// fmt.Printf("%+v\n", []interface{}{a, b})
	fmt.Printf("%+v\n", []interface{}{jump([]int{2, 3, 1, 1, 4})})
	// fmt.Printf("%+v\n", []interface{}{jump([]int{9, 8, 2, 2, 0, 2, 2, 0, 4, 1, 5, 7, 9, 6, 6, 0, 6, 5, 0, 5})})

}

func jump(nums []int) int {
	step := 0
	for i := 0; i < len(nums)-1; step++ {
		i = rangeMax(nums[i], nums, i+1)
	}
	return step
}

func rangeMax(size int, nums []int, from int) int {
	maxK := from
	maxV := 1
	for i := from; i < from+size && i < len(nums); i++ {
		if i == len(nums)-1 {
			return i
		}
		if i+nums[i] > maxV {
			maxK = i
			maxV = nums[i] + i
		}
	}
	return maxK
}
