package main

import (
	"fmt"
)

func validMountainArray(A []int) bool {
	up := true
	if len(A) < 3 {
		return false
	}
	if A[0] > A[1] {
		return false
	}
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		}
		if up && A[i] < A[i-1] {
			up = false
			continue
		}
		if up && A[i] < A[i-1] {
			return false
		}
		if !up && A[i] > A[i-1] {
			return false
		}
	}
	return !up
}
func main() {
	fmt.Printf("%+v\n", []interface{}{validMountainArray([]int{3, 5, 5})})
}
