package main

import (
	"fmt"
)

func reverseWords(s string) string {
	bytes := []byte(s)
	i, j, space := 0, 0, 0
	l := len(bytes)
	for j < l {
		if bytes[j] != byte(' ') {
			j++
		} else {
			space = j
			for i < j {
				bytes[i], bytes[j-1] = bytes[j-1], bytes[i]
				i++
				j--
			}
			i, j = space+1, space+1
		}
	}
	for i < j {
		bytes[i], bytes[j-1] = bytes[j-1], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func main() {
	fmt.Printf("%+v\n", []interface{}{reverseWords("hello world")})
}
