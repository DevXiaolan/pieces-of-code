package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	m := map[string]byte{}
	rm := map[byte]string{}
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	for i, s := range strs {
		_, ok := m[s]
		_, rok := rm[byte(pattern[i])]
		if !ok && !rok {
			m[s] = pattern[i]
			rm[byte(pattern[i])] = s
		} else {
			if !(m[s] == pattern[i] && rm[byte(pattern[i])] == s) {
				return false
			}
		}
	}

	return true
}
func main() {
	fmt.Printf("%+v\n", []interface{}{wordPattern("abba", "dog dat dat dog")})
}
