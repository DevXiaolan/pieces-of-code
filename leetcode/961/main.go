package main

import (
	"fmt"
)

func repeatedNTimes(A []int) int {
	n := len(A) / 2
	count := [10000]int{}
	for _, v := range A {
		count[v]++
		if count[v] == n {
			return v
		}
	}
	return 0
}

func main() {
	fmt.Printf("%+v\n", []interface{}{repeatedNTimes([]int{1, 2, 3, 3})})
}
