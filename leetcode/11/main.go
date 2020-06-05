package main

import "fmt"

func maxArea(height []int) int {
	start, end, max := 0, len(height)-1, 0
	for start < end {
		tmp := 0
		if height[start] > height[end] {
			tmp = height[end] * (end - start)
			end--
		} else {
			tmp = height[start] * (end - start)
			start++
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}
