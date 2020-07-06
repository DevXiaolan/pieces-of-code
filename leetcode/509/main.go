package main

import (
	"fmt"
)

var M map[int]int

func init() {
	M = map[int]int{}
}

func fib(N int) int {
	if N == 0 {
		M[N] = 0
		return 0
	}
	if N == 1 {
		M[N] = 1
		return 1
	}
	ans := fib(N-2) + fib(N-1)
	M[N] = ans
	return ans
}

func main() {
	fmt.Printf("%+v\n", []interface{}{fib(4)})
}
