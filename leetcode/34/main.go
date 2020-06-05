package main

import "fmt"

func main() {
	n := searchRange([]int{1}, 1)
	fmt.Printf("%+v\n", []interface{}{n})
}
func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	t := []int{}
	for i := 0; i < numsLen; i++ {
		if nums[i] == target {
			t = append(t, i)
		} else if len(t) > 0 {
			break
		}
	}
	if len(t) > 0 {
		return []int{t[0], t[len(t)-1]}
	}
	return []int{-1, -1}
}
