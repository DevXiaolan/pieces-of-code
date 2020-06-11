package main

import (
	"fmt"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	n := minutesToTest/minutesToDie + 1
	i := 0
	sum := 1
	for {
		i++
		sum *= n
		if sum >= buckets {
			break
		}
	}
	return i
}

func main() {
	fmt.Printf("%+v\n", []interface{}{poorPigs(1000, 15, 60)})
}
