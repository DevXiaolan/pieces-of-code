package main

import (
	"fmt"
)

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 2
	}
	setIn, start, end := 0, 0, 0
	for end < len(chars)-1 {
		if chars[end] == chars[end+1] {
			end++
		} else {
			if end > start {
				chars[setIn] = chars[start]
				offset := end - start
				bits := []int{}
				for offset > 0 {
					bits = append(bits, offset%10)
					offset = offset / 10
				}
				for i := len(bits) - 1; i >= 0; i-- {
					chars[setIn+1] = byte(bits[i] + 49)
					setIn++
				}
			}
			setIn++
			start = end + 1
			end = start
		}
	}
	chars[setIn] = chars[start]
	offset := end - start + 1

	bits := []int{}
	for offset > 0 {
		bits = append(bits, offset%10)
		offset = offset / 10
	}

	for i := len(bits) - 1; i >= 0; i-- {
		chars[setIn+1] = byte(bits[i] + 48)
		setIn++
	}
	return setIn + 1
}

func main() {
	fmt.Printf("%+v\n", []interface{}{compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})})
}
