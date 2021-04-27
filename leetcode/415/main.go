package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {

	x, y := len(num1)-1, len(num2)-1
	var b byte = 0
	var result []byte
	if x > y {
		result = make([]byte, x+2)
	} else {
		result = make([]byte, y+2)
	}
	index := len(result) - 1
	for x >= 0 || y >= 0 {
		var tmp byte = b
		if x >= 0 {
			tmp += num1[x] - '0'
		}
		if y >= 0 {
			tmp += num2[y] - '0'
		}
		if tmp > 9 {
			b = 1
			tmp = tmp % 10
		} else {
			b = 0
		}
		result[index] = tmp + '0'
		x--
		y--
		index--
	}
	if b > 0 {
		result[0] = '1'
	}
	if result[0] == 0 {
		result = result[1:]
	}
	// fmt.Printf("%+v\n", []interface{}{result})
	return string(result)
}

func main() {
	fmt.Printf("%+v\n", []interface{}{addStrings("9876", "321")})
}
