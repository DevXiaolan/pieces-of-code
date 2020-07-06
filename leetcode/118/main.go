package main

import (
	"fmt"
)

func nextTick(tick []int) []int {
	l := len(tick)
	next := make([]int, l+1)
	for k := range next {
		if k == 0 || k == l {
			next[k] = 1
		} else {
			next[k] = tick[k-1] + tick[k]
		}
	}
	return next
}

func generate(numRows int) [][]int {
	result := [][]int{}
	tick := []int{1}
	for i := 0; i < numRows; i++ {
		result = append(result, tick)
		tick = nextTick(tick)
	}

	return result
}

func main() {
	fmt.Printf("%+v\n", []interface{}{generate(5)})
}
