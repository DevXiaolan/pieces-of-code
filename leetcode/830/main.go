package main

import (
	"fmt"
)

func largeGroupPositions(s string) [][]int {
	start, end := 0, 0
	result := [][]int{}
	for k := 0; k < len(s); k++ {
		if k == len(s)-1 || s[k] != s[k+1] {
			if end-start >= 2 {
				// large
				result = append(result, []int{start, end})
			}
			start, end = k+1, k+1
		} else {
			end++
		}
	}
	return result
}

func main() {
	// fmt.Printf("%+v\n", []interface{}{largeGroupPositions("abcdddeeeeaabbbcd")})
	fmt.Printf("%+v\n", []interface{}{largeGroupPositions("aaa")})
}
