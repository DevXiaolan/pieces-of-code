package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%+v\n", []interface{}{dayOfYear("2004-03-01")})
}

func isR(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%4 == 0 && year%100 != 0 {
		return true
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	return false
}

func parse(date string) (int, int, int) {
	pieces := strings.Split(date, "-")
	year, _ := strconv.Atoi(pieces[0])
	month, _ := strconv.Atoi(pieces[1])
	day, _ := strconv.Atoi(pieces[2])
	return year, month, day
}

func dayOfYear(date string) int {
	cache := [][]int{
		[]int{0, 0},
		[]int{31, 31},
		[]int{59, 60},
		[]int{90, 91},
		[]int{120, 121},
		[]int{151, 152},
		[]int{181, 182},
		[]int{212, 213},
		[]int{243, 244},
		[]int{273, 274},
		[]int{304, 305},
		[]int{334, 335},
	}
	year, month, day := parse(date)
	result := day
	if isR(year) {
		result += cache[month-1][1]
	} else {
		result += cache[month-1][0]
	}
	return result
}
