package main

import "fmt"

func maximumSwap(num int) int {
	arr := []int{}
	for num > 0 {
		t := num % 10
		num = num / 10
		arr = append(arr, t)
	}

	swap := true
	for i := len(arr) - 1; i > 0 && swap; i-- {
		maxI := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] >= arr[maxI] {
				maxI = j
			}
		}
		if i != maxI && arr[maxI] != arr[i] {
			arr[maxI], arr[i] = arr[i], arr[maxI]
			swap = false
		}
	}
	sum := 0
	for i := len(arr) - 1; i >= 0; i-- {

		sum = sum*10 + arr[i]
	}
	return sum
}

func main() {
	fmt.Printf("%+v\n", []interface{}{maximumSwap(98368)})
}
