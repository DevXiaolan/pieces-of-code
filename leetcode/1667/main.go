package main

import (
	"fmt"
)

func maxAliveYear(birth []int, death []int) int {
	years := make([]int, 102, 102)
	for i := 0; i < len(birth); i++ {
		years[birth[i]-1900]++
		years[death[i]-1900+1]--
	}

	sum := 0
	max := 0
	year := 0
	for i := 0; i < len(years); i++ {
		sum += years[i]
		if sum > max {
			max = sum
			year = i + 1900
		}
	}
	return year
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxAliveYear([]int{1900, 1901, 1950}, []int{1948, 1951, 2000})})
}
