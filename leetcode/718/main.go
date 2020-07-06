package main

import (
	"fmt"
)

func findLength(A []int, B []int) int {
	i, j, max, tmp := 0, 0, 0, 0
	for i < len(A) {
		for j < len(B) && i < len(A) {
			if A[i] == B[j] {
				tmp++
				if tmp > max {
					max = tmp
				}
				fmt.Printf("%+v\n", []interface{}{i, j, max})
				i++
			} else if tmp > 0 {
				break
			}
			j++
		}
		i++
		j = 0
		tmp = 0
	}
	return max
}
func main() {
	fmt.Printf("%+v\n", []interface{}{findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1})})
	// fmt.Printf("%+v\n", []interface{}{findLength([]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0})})
	// fmt.Printf("%+v\n", []interface{}{findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})})
}
