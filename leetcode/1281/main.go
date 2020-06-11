package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	p := 1
	s := 0
	for n > 0 {
		bit := n % 10
		n = n / 10
		p *= bit
		s += bit
	}
	return p - s
}

func main() {
	fmt.Printf("%+v\n", []interface{}{subtractProductAndSum(234)})
}
