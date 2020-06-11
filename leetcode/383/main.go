package main

import (
	"fmt"
)

func find(inputs []byte, search byte, offset int) int {
	l := len(inputs)
	for i := offset; i < l; i++ {
		if inputs[i] == search {
			return i
		}
	}
	return -1
}

func canConstruct(ransomNote string, magazine string) bool {
	jump := [26]int{}
	l := len(ransomNote)
	bytes := []byte(magazine)
	for i := 0; i < l; i++ {
		search := byte(ransomNote[i])
		index := find(bytes, search, jump[search-97])
		if index == -1 {
			return false
		}
		jump[search-97] = index + 1
	}
	return true
}

func main() {
	fmt.Printf("%+v\n", []interface{}{canConstruct("aa", "aab")})
}
