package main

import (
	"fmt"
)

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	return addDigits(num%10 + addDigits(num/10))
}

func main() {
	fmt.Printf("%+v\n", []interface{}{addDigits(1234)})
}
