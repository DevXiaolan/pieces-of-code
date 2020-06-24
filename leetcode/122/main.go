package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit, buy := 0, -1
	for k := 0; k < len(prices); k++ {
		v := prices[k]
		if k == len(prices)-1 {
			if buy != -1 {
				profit += v - buy
			}
			break
		}

		if v > prices[k+1] && (buy != -1) {
			profit += v - buy
			buy = -1
		}
		if v < prices[k+1] {
			if buy == -1 {
				buy = v
			}
		}
	}
	return profit
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxProfit([]int{2, 1, 2, 0, 1})})
}
