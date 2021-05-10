package main

import "fmt"

 
func findRepeatNumber(nums []int) int {
	bits := make([]int, len(nums))
	for _,v := range nums {
		bits[v]++
		if bits[v] >1 {
			return v
		}
	}
	return 0
}
func main(){
	fmt.Printf("%+v\n",[]interface{}{findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})})
}