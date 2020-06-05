package main

import "fmt"

func main() {

	fmt.Printf("%+v\n", []interface{}{permute([]int{1, 2, 3})})
}

func permute(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{nums}
	}
	if len(nums) == 2 {
		return [][]int{
			[]int{nums[0], nums[1]},
			[]int{nums[1], nums[0]},
		}
	}
	tmp := [][]int{}
	for k, num := range nums {
		rest := []int{}
		rest = append(rest, nums[0:k]...)
		rest = append(rest, nums[k+1:]...)
		for _, item := range permute(rest) {
			tmp = append(tmp, append(item, num))
		}
	}
	return tmp
}
