package main

import (
	"fmt"
)

var dict map[byte][]byte

func init() {
	dict = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	sl1 := dict[digits[0]]
	if len(digits) == 1 {
		for _, v := range sl1 {
			result = append(result, string(v))
		}
		return result
	}
	nexts := letterCombinations(digits[1:])
	for _, v := range sl1 {
		for _, str := range nexts {
			result = append(result, string(v)+str)
		}
	}
	return result
}

func main() {
	fmt.Printf("%+v\n", []interface{}{letterCombinations("23")})
}
