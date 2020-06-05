package main

import "fmt"

const MAX, MIN = int(^uint32(0) >> 1), ^int(^uint32(0) >> 1)

func threeSumClosest(nums []int, target int) int {
	sorted := quickSort(nums)
	fmt.Printf("%+v\n", sorted)
	closest := sorted[0] + sorted[1] + sorted[len(sorted)-1]
	for i := 0; i < len(sorted)-2; i++ {
		fmt.Printf("Turn %+v\n", i)
		start := i + 1
		end := len(sorted) - 1
		//closest = sorted[i] + sorted[start] + sorted[end]
		for start < end {
			fmt.Printf("before %+v\n", []int{i, start, end})
			if sorted[i]+sorted[start]+sorted[end] > target {

				if (target-sorted[i]-sorted[start]-sorted[end])*(target-sorted[i]-sorted[start]-sorted[end]) < (target-closest)*(target-closest) {
					fmt.Printf("Got %+v\n", sorted[i]+sorted[start]+sorted[end])
					closest = sorted[i] + sorted[start] + sorted[end]
				}
				start++
			} else if sorted[i]+sorted[start]+sorted[end] < target {

				if (target-sorted[i]-sorted[start]-sorted[end])*(target-sorted[i]-sorted[start]-sorted[end]) <= (target-closest)*(target-closest) {

					fmt.Printf("Got %+v\n", sorted[i]+sorted[start]+sorted[end])
					closest = sorted[i] + sorted[start] + sorted[end]
				}
				end--
			} else if sorted[i]+sorted[start]+sorted[end] == target {
				closest = target
				break
			}

		}
	}
	return closest
}

func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	left, rigth := []int{}, []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			left = append(left, nums[i])
		} else if nums[i] <= nums[0] {
			rigth = append(rigth, nums[i])
		}
	}
	return append(append(quickSort(left), nums[0]), quickSort(rigth)...)
}

func main() {
	// fmt.Printf("%+v\n", quickSort([]int{-1, 2, 1, -4}))
	fmt.Printf("Final %+v\n", threeSumClosest([]int{0, 2, 1, -3}, 1))
}
