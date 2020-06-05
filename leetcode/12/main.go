package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	RomanList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	IntList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := []string{}
	for i := 0; i < 13; i++ {
		for IntList[i] <= num {
			num -= IntList[i]
			result = append(result, RomanList[i])
		}
	}
	return strings.Join(result, "")
}

func main() {
	// MCMXCIV
	fmt.Println(intToRoman(1994))
}
