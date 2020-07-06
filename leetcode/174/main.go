package main

import (
	"fmt"
)

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp, dpHp := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dpHp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				if dungeon[i][j] > 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 - dungeon[i][j]
				}
				dpHp[i][j] = dp[i][j] + dungeon[i][j]
			} else if j == 0 {
				dpHp[i][j] = dpHp[i-1][j] + dungeon[i][j]
				if dpHp[i][j] > 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					// dp[i][j] = dp[i-1][j] - dungeon[i][j]
					dp[i][j] = dp[i-1][j] - (dpHp[i-1][j] + dungeon[i][j] - 1)
					dpHp[i][j] = 1
				}
			} else if i == 0 {
				dpHp[i][j] = dpHp[i][j-1] + dungeon[i][j]
				if dpHp[i][j] > 0 {
					dp[i][j] = dp[i][j-1]
				} else {
					// dp[i][j] = dp[i][j-1] - dungeon[i][j]
					dp[i][j] = dp[i][j-1] - (dpHp[i][j-1] + dungeon[i][j] - 1)
					dpHp[i][j] = 1
				}
			} else {
				// i-1
				var sl1, sl2, hp1, hp2 int
				if dpHp[i-1][j]+dungeon[i][j] > 0 {
					sl1 = dp[i-1][j]
					hp1 = dpHp[i-1][j] + dungeon[i][j]
				} else {
					sl1 = dp[i-1][j] - (dpHp[i-1][j] + dungeon[i][j] - 1)
					hp1 = 1
				}
				if dpHp[i][j-1]+dungeon[i][j] > 0 {
					sl2 = dp[i][j-1]
					hp2 = dpHp[i][j-1] + dungeon[i][j]
				} else {
					sl2 = dp[i][j-1] - (dpHp[i][j-1] + dungeon[i][j] - 1)
					hp2 = 1
				}
				if sl1 < sl2 {
					dp[i][j] = sl1
					dpHp[i][j] = hp1
				} else {
					dp[i][j] = sl2
					dpHp[i][j] = hp2
				}
			}
		}
	}
	// fmt.Printf("%+v\n", []interface{}{dp, dpHp})
	return dp[m-1][n-1]
}

func main() {
	area := [][]int{
		{1},
		{-2},
		{1},
	}
	fmt.Printf("%+v\n", []interface{}{calculateMinimumHP(area)})
}
