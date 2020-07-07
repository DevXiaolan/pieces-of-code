package main

import (
	"fmt"
)

func fix(a int) int {
	if a < 0 {
		return 0
	}
	return a
}

func maxSubArray(nums []int) int {
	// dp[i] 表示以i为右端点的连续最大和
	// dp[i] = dp[i-1] + nums[i]
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := make([]int, l)
	dp[0] = nums[0]
	if l == 1 {
		return dp[0]
	}
	max := nums[0]
	for i := 1; i < l; i++ {
		dp[i] = fix(dp[i-1]) + nums[i]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Printf("%+v\n", []interface{}{maxSubArray([]int{1, 2})})
}
