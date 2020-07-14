package main

import (
	"fmt"
)

func min(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
func minimumTotal(triangle [][]int) int {
	// dp[i][j]  表示从顶点到当前点的最小路径
	// dp[i][j] = triangle[i][j] + min(dp[i-1][j-1], dp[i-1][j])
	for i, layer := range triangle {
		if i == 0 {
			continue
		}
		for j := range layer {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(layer)-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += min([]int{triangle[i-1][j-1], triangle[i-1][j]})
			}
		}
	}
	// fmt.Printf("%+v\n", []interface{}{triangle})
	return min(triangle[len(triangle)-1])
}
func main() {
	fmt.Printf("%+v\n", []interface{}{minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	})})
}
