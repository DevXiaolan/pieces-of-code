package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	cp1 := make([]int, m)
	copy(cp1, nums1)
	p1, p2, index := 0, 0, 0
	for p1 < m || p2 < n {
		fmt.Printf("%+v\n", []interface{}{cp1[p1], nums2[p2]})
		if p1 >= m {
			nums1[index] = nums2[p2]
			index++
			p2++
		} else if p2 >= n {
			nums1[index] = cp1[p1]
			index++
			p1++
		} else {
			if cp1[p1] < nums2[p2] {
				nums1[index] = cp1[p1]
				index++
				p1++
			} else {
				nums1[index] = nums2[p2]
				index++
				p2++
			}
		}
		fmt.Printf("%+v\n", []interface{}{nums1})
	}
}

func main() {
	arr1 := []int{4, 5, 6, 0, 0, 0}
	merge(arr1, 3, []int{1, 2, 3}, 3)
	fmt.Printf("%+v\n", []interface{}{arr1})
}
