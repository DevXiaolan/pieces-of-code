package main

import "fmt"

func main() {
	n := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Printf("%+v\n", []interface{}{n})
}

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	step := numsLen / 2
	for i := numsLen / 2; i >= 0 && i < numsLen; {
		step = step / 2
		if step == 0 {
			step = 1
		}
		//fmt.Printf("%+v\n", []interface{}{i, step})
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			if i == 0 {
				return 0
			}
			if nums[i-1] < target {
				return i
			}
			i = i - step
		} else {
			if i+1 > numsLen {
				return i + 1
			} else if nums[i+1] > target {
				return i + 1
			}
			i = i + step
		}
	}
	return 0
}
