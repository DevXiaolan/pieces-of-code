package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strs := []string{}
	l := len(s) - 1
	i, j := l, l
	for i >= 0 {
		if s[i] != ' ' {
			i--
		} else {
			if i == j {
				i--
				j--
				continue
			} else {
				strs = append(strs, s[i+1:j+1])
				i--
				j = i
			}
		}
	}
	if i != j {
		strs = append(strs, s[:j+1])
	}
	return strings.Join(strs, " ")
}

func main() {
	fmt.Printf("%+v\n", []interface{}{reverseWords("the sky is blue")})
}
