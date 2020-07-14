package main

import (
	"fmt"
)

func findMaxMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

func subSort(array []int) []int {
	if len(array) < 2 {
		return []int{-1, -1}
	}
	start, end := 0, len(array)-1
	for start < len(array)-1 && array[start] < array[start+1] {
		start++
	}
	for end > 0 && array[end] > array[end-1] {
		end--
	}
	if end == 0 {
		return []int{-1, -1}
	}

	max, min := findMaxMin(array[start+1 : end])
	for end < len(array)-1 && max > array[end] {
		max = array[end]
		end++
	}
	for start > 0 && min < array[start] {
		min = array[start]
		start--
	}
	return []int{start + 1, end}
}
func main() {
	//
	// fmt.Printf("%+v\n", []interface{}{subSort([]int{1, 3, 9, 7, 5})})
	fmt.Printf("%+v\n", []interface{}{subSort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})})
}
