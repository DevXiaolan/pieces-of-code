package main

import "fmt"

func romanToInt(s string) int {

	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	result := 0
	startAt := 0
	for len(s) > startAt {
		index := 0

		for {
			if startAt+index >= len(s) {
				break
			}
			if _, isPresent := m[s[startAt:startAt+index+1]]; !isPresent {
				break
			}
			index++

		}

		result += m[s[startAt:startAt+index]]
		startAt += index
	}
	return result
}

func main() {

	// 1994
	fmt.Println(romanToInt("MCMXCIV"))
}
