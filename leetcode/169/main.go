package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			major = v
			count = 1
		} else if v == major {
			count++
		} else if v != major {
			count--
		}
	}
	return major
}

func main() {
	fmt.Printf("%+v\n", []interface{}{majorityElement([]int{2, 2, 1, 1, 1, 2, 2})})
}
