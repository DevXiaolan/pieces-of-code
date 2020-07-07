package main

import (
	"fmt"
)

func minV(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairs(cost []int) int {
	// dp[i] = min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	l := len(cost)
	dp := make([]int, l+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < l+1; i++ {
		dp[i] = minV(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[l]
}

func main() {
	fmt.Printf("%+v\n", []interface{}{minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})})
	// fmt.Printf("%+v\n", []interface{}{minCostClimbingStairs([]int{10, 15, 20})})
}
