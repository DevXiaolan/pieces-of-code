package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	s := 0
	for _,v := range nums {
		s = s^v
	}
	return s
}
func main() {
	fmt.Printf("%+v\n", []interface{}{singleNumber([]int{4, 1, 2, 1, 2})})
}

/**
这题太简单了，单身狗解法，全部异或，剩谁谁尴尬
*/