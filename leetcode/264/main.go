package main

import (
	"fmt"
)

var nums []int
var m map[int]bool

func init() {
	nums = []int{1}
	count := 0
	m = map[int]bool{}
	i2, i3, i5 := 0, 0, 0
	for count < 1691 { // 1691
		n2, n3, n5 := nums[i2]*2, nums[i3]*3, nums[i5]*5
		minNum := min(n2, n3, n5)
		switch minNum {
		case n2:
			if !m[n2] {
				nums = append(nums, n2)
				count++
				m[n2] = true
			}
			i2++
		case n3:
			if !m[n3] {
				nums = append(nums, n3)
				count++
				m[n3] = true
			}
			i3++
		case n5:
			if !m[n5] {
				nums = append(nums, n5)
				count++
				m[n5] = true
			}
			i5++
		}
	}
}

func min(a int, b int, c int) int {
	var tmp int = a
	if tmp > b {
		tmp = b
	}
	if tmp > c {
		tmp = c
	}
	return tmp
}

func nthUglyNumber(n int) int {
	return nums[n-1]
}

func main() {
	fmt.Printf("%+v\n", []interface{}{nthUglyNumber(10)})
}
