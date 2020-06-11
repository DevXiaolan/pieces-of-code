package main

import (
	"fmt"
)

func turn(A []int)[]int{
	l := len(A)
	i,j := 0,l-1
	for i<=j {
		A[i],A[j] = 1^A[j],1^A[i]
		i++
		j--
	}
	return A
}

func flipAndInvertImage(A [][]int) [][]int {
	l := len(A)
	for i:=0;i<l;i++ {
		turn(A[i])
	}
	return A
}

func main() {
	fmt.Printf("%+v\n", []interface{}{flipAndInvertImage([][]int{
		[]int{1, 1, 0},
		[]int{1, 0, 1},
		[]int{0, 0, 0},
	})})
}
