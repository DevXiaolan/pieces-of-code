package main

import (
	"fmt"
)

func uniqueOccurrences(arr []int) bool {
	max, m := 0, map[int]int{}
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
		if m[v] > max {
			max = m[v]
		}
	}

	check := make([]int, max)

	for _, v := range m {
		if check[v-1] == 0 {
			check[v-1] = 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%+v\n", []interface{}{uniqueOccurrences([]int{1, 2})})
}
