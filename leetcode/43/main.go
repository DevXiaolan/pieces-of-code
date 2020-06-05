package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", []interface{}{multiply("123", "456")})
}

func multiply(num1 string, num2 string) string {

	return ""
}

func arrFy(str string) []int {
	l := len(str)
	arr := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		arr = append(arr, str[i])
	}
}
