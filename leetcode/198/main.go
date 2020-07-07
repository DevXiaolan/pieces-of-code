package main

import (
	"fmt"
)

func rob(nums []int) int {
	// dp[i] 表示从i出发
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
		return max(dp[0], dp[1])
	}
	dp[l-3] = nums[l-3] + dp[l-1]
	if l == 3 {
		return max(dp[0], dp[1])
	}
	for i := l - 4; i >= 0; i-- {
		dp[i] = max(dp[i+2], dp[i+3]) + nums[i]
	}
	return max(dp[0], dp[1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 12
	fmt.Printf("%+v\n", []interface{}{rob([]int{2, 7, 9, 3, 1})})
}
