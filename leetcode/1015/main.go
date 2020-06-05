package main

import (
	"fmt"
)

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	count := 1
	var i = 1
	for {
		if i%K == 0 {
			return count
		}
		i = i % K
		i = i*10 + 1
		count++
	}
}

func main() {
	fmt.Printf("%+v\n", []interface{}{smallestRepunitDivByK(7)})
}
