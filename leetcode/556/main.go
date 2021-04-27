package main

import "fmt"

func sum(arr []int) int {
	l := len(arr) - 1
	s := 0
	for i := 0; i <= l; i++ {
		s = s*10 + arr[l-i]
	}
	return s
}
func nextGreaterElement(n int) int {
	arr := []int{}
	for n > 0 {
		arr = append(arr, n%10)
		n = n / 10
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				// find
				arr[i], arr[j] = arr[j], arr[i]
				return sum(arr)
			}
		}
	}
	return -1
}
func main() {
	fmt.Printf("%+v\n", []interface{}{nextGreaterElement(312)})
}
