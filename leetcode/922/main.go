package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", []interface{}{sort([]int{3, 4, 5, 6, 7, 8})})
}

func sort(arr []int) []int {
	length := len(arr)
	iOdd := 1
	for i := 0; i < length-1; {
		if arr[i]%2 == 0 {
			i += 2
			continue
		} else {
			arr[i], arr[iOdd] = arr[iOdd], arr[i]
			iOdd += 2
		}
	}
	return arr
}
