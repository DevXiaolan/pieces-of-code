package main

import (
	"fmt"
)

func waysToStep(n int) int {
	// dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	if n == 0 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	if n == 1 {
		return dp[n-1]
	}
	dp[1] = dp[0] + 1
	if n == 2 {
		return dp[n-1]
	}
	dp[2] = dp[0] + dp[1] + 1
	if n == 3 {
		return dp[n-1]
	}
	for i := 3; i < n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}
	return dp[n-1]
}

func main() {
	fmt.Printf("%+v\n", []interface{}{waysToStep(3)})
}
