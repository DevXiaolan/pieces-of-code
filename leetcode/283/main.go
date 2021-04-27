package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	i, j, l := 0, 0, len(nums)
	for i < l {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for j < l {
		nums[j] = 0
		j++
	}
	fmt.Printf("%+v\n", []interface{}{nums, i, j})
}
func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})

}
