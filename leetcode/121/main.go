package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit, buy := 0, 2<<32
	for _, v := range prices {
		if v < buy {
			buy = v
		}
		if v-buy > profit {
			profit = v - buy
		}
	}
	return profit
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxProfit([]int{7, 1, 5, 3, 6, 4})})
}
