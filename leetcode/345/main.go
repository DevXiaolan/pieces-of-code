package main

import (
	"fmt"
)

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	slice := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isVowel(s[i]) {
			i++
		} else if !isVowel(s[j]) {
			j--
		} else {
			slice[i] = slice[i] ^ slice[j]
			slice[j] = slice[i] ^ slice[j]
			slice[i] = slice[i] ^ slice[j]
			i++
			j--
		}
	}
	return string(slice)
}

func main() {
	str := "leetcode"
	fmt.Printf("%+v\n", []interface{}{reverseVowels(str)})
}
