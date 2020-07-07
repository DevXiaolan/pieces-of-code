package main

import (
	"fmt"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func massage(nums []int) int {
	// dp[i] 表示从开始接单的最大收益
	// dp[i] = max(dp[i+2],dp[i+3]) + nums[i]
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	dp[l-1] = nums[l-1]
	if l == 1 {
		return dp[l-1]
	}
	dp[l-2] = nums[l-2]
	if l == 2 {
		return max(dp[l-1], dp[l-2])
	}
	dp[l-3] = max(dp[l-1]+nums[l-3], dp[l-2])
	if l == 3 {
		return dp[0]
	}
	for i := l - 4; i >= 0; i-- {
		dp[i] = max(dp[i+2], dp[i+3]) + nums[i]
	}
	// fmt.Printf("%+v\n", []interface{}{dp})
	return max(dp[0], dp[1])
}

func main() {
	fmt.Printf("%+v\n", []interface{}{massage([]int{1, 2, 3, 1})})
}
