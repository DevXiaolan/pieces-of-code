package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func main() {
	fmt.Printf("%+v\n", []interface{}{isPowerOfTwo(30)})
}
