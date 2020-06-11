package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	l := len(arr)
	var count [501]int
	for i := 0; i < l; i++ {
		count[arr[i]]++
	}
	for i:=500;i>0;i-- {
		if i == count[i] {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Printf("%+v\n", []interface{}{findLucky([]int{1, 2, 2, 3, 4})})
}
