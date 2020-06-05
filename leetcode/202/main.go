package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", []interface{}{happy(19)})
}

func happy(a int) bool {
	cache := make(map[int]bool)
	tmp := a
	for tmp != 1 {
		sum := 0
		for tmp > 0 {
			sum += (tmp % 10) * (tmp % 10)
			tmp /= 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := cache[sum]; ok {
			return false
		}
		cache[sum] = true
		tmp = sum
	}
	return true
}
