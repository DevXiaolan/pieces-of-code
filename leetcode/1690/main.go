package main

import (
	"fmt"
)

func isFlipedString(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return false
	}
	if s1 == "" {
		return true
	}
	// 找对齐位
	ds2 := s2 + s2
	// fmt.Printf("%+v\n", []interface{}{ds2})
	count := 0
	j := 0
	for i := 0; i < 2*l2; i++ {
		if ds2[i] != s1[j] {
			count = 0
			j = 0
		} else {
			j++

			count++
			if count == l1 {
				return true
			}
			if j >= l1 {
				return false
			}
		}
	}
	return false
}

func main() {
	fmt.Printf("%+v\n", []interface{}{isFlipedString("aba", "bab")})
}
