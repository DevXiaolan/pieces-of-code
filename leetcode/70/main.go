package main

import (
	"fmt"
)

var M map[int]int

func init() {
	M = map[int]int{}
}
func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		M[n] = n
		return n
	}
	if _, ok := M[n]; ok {
		return M[n]
	}
	ans := climbStairs(n-2) + climbStairs(n-1)
	M[n] = ans
	return ans
}

func main() {
	fmt.Printf("%+v\n", []interface{}{climbStairs(3)})
}
