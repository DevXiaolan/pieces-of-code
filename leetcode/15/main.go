package main

import "fmt"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sorted := quicksort(nums)
	// fmt.Printf("%+v\n", sorted)
	if len(sorted) == 0 || sorted[0] < 0 || sorted[len(sorted)-1] > 0 {
		return result
	}
	for i := 0; i < len(sorted)-2; i++ {

		start := i + 1
		end := len(sorted) - 1
		// fmt.Printf("%+v,%+v,%+v\n", i, start, end)
		for start < end {
			if sorted[i]+sorted[start]+sorted[end] > 0 {
				start++
			} else if sorted[i]+sorted[start]+sorted[end] < 0 {
				end--
			}
			//fmt.Printf("%+v,%+v,%+v\n", i, start, end)
			if start < end && sorted[i]+sorted[start]+sorted[end] == 0 {
				// start++
				result = append(result, []int{sorted[i], sorted[start], sorted[end]})
				// fmt.Printf("%+v,%+v,%+v\n", i, start, end)
				for start < end && sorted[start+1] == sorted[start] {
					start++
				}
				start++
				// break
			}

		}

		// if start < end && sorted[i]+sorted[start]+sorted[end] == 0 {
		// 	fmt.Printf("%+v\n", []int{sorted[i], sorted[start], sorted[end]})
		// 	result = append(result, []int{sorted[i], sorted[start], sorted[end]})
		// }
		// break
		for i < len(sorted)-2 && sorted[i+1] == sorted[i] {
			i++
		}
	}
	return result
}

func quicksort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}
	left, right := []int{}, []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	return append(append(quicksort(left), nums[0]), quicksort(right)...)
}

func main() {
	fmt.Printf("%+v\n", threeSum([]int{0, 0, 0}))
	//fmt.Printf("%+v\n", quicksort([]int{-1, 0, 1, 2, -1, -4}))
}
