package main

import "fmt"

func checkPerfectNumber(num int) bool {
	sum := 1
	head := num
	for i := 2; i < head; i++ {
		if num%i == 0 {
			sum = sum + i + num/i
			head = head / i
		}
		if sum > num {
			return false
		}
	}

	return sum == num
}

func main() {
	fmt.Printf("%+v\n", []interface{}{checkPerfectNumber(6)})
}
