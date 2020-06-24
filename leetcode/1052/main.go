package main

import (
	"fmt"
)

func maxSatisfied(customers []int, grumpy []int, X int) int {
	l := len(customers)
	maxIdx,maxWaste,satisfy := 0,0,0

	for i:=0;i<l;i++ {
		if grumpy[i]==0 {
			satisfy += customers[i]
		}
		waste := 0
		for j:=0;j<X&&(i+j)<l;j++ {
			waste += customers[i+j]*grumpy[i+j]
		}
		if waste>maxWaste {
			maxWaste = waste
			maxIdx = i
		}
	}
	for i:=0;i<X;i++ {
		if grumpy[maxIdx+i] == 1 {
			satisfy += customers[maxIdx+i]
		}
	}
	return satisfy
}
func main() {
	fmt.Printf("%+v\n", []interface{}{maxSatisfied(
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1},
		3,
	),
	},
	)
}
