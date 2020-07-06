package main

import (
	"fmt"
)

func isC(b byte) bool {
	if (b > 64 && b < 91) || (b > 96 && b < 123) {
		return true
	}
	return false
}
func reverseOnlyLetters(S string) string {
	bytes := []byte(S)
	i, j := 0, len(S)-1
	for i < j {
		for i < j && !isC(bytes[i]) {
			i++
		}
		for j >= 0 && !isC(bytes[j]) {
			j--
		}

		if i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		}
	}
	return string(bytes)
}

func main() {
	fmt.Printf("%+v\n", []interface{}{reverseOnlyLetters("7_28]")})
}
