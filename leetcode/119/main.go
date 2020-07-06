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

func getRow(rowIndex int) []int {
	tick := []int{1}
	for i := 0; i < rowIndex; i++ {
		tick = nextTick(tick)
	}
	return tick
}

func main() {
	fmt.Printf("%+v\n", []interface{}{getRow(3)})
}
