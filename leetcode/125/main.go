package main

import (
	"fmt"
)

func isBaz(b byte) bool {
	if (b >= 97 && b <= 122) || (b >= 48 && b <= 57) {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := len(s)
	i, j := 0, l-1
	for i < j {
		if !isBaz(s[i]) {
			i++
			continue
		}
		if !isBaz(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func main() {
	fmt.Printf("%+v\n", []interface{}{isPalindrome("race a car")})
}
