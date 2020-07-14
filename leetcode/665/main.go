package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] <= nums[i] {
			continue
		}
		count++
		if i-2 >= 0 && nums[i-2] > nums[i] { //这种情况是突然遇到了一个比前一个数小的数字，同时小于它前一个的前一个数字
			nums[i] = nums[i-1] //这句代码是更新第i个位置的数据，保证其非递减行，但不要担心，因为在这之前，计数器已经+1了
		} else {
			nums[i-1] = nums[i]
		}

	}
	return count <= 1
}
func main() {
	fmt.Printf("%+v\n", []interface{}{checkPossibility([]int{2, 3, 3, 2, 4})})
}
