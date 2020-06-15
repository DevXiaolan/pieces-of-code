package main

import (
	"fmt"
)
func duplicateZeros(arr []int)  {
	l := len(arr)
	full := make([]int, 2*l)
	k:=0
	z:=0
	for _,v := range arr {
		full[k] = v
		k++
		if v==0 {
			z++
			full[k] = v
			k++
		}
		
	}
	copy(arr,full)
}

func main() {
	arr := []int{1,0,2,3,0,4,5,0}
	duplicateZeros(arr)
	fmt.Printf("%+v\n",[]interface{}{arr})
}