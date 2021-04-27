package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		sum := numbers[start] + numbers[end]
		if sum == target {
			return []int{start + 1, end + 1}
		}
		if sum > target {
			end--
		} else {
			start++
		}
	}
	return []int{0, 0}
}

func main() {
	fmt.Printf("%+v\n", []interface{}{twoSum([]int{2, 7, 11, 15}, 9)})
}
