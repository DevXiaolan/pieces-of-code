package main

import "fmt"

func partitionDisjoint(A []int) int {
	fmt.Printf("%+v\n",[]interface{}{A})
	l := len(A)
	R := make([]int,l)
	R[l-1] = A[l-1]
	for k:=l-2;k>=0;k-- {
		if A[k]<R[k+1] {
			R[k] = A[k]
		}else{
			R[k] = R[k+1]
		}
	}
	fmt.Printf("%+v\n",[]interface{}{R})
	LMAX := A[0]
	k:=0
	for k=0;k<l-1;k++{
		if A[k]>= LMAX {
			LMAX = A[k]
		}
		if LMAX <= R[k+1] {
			return k+1
		}
	}
	return k
}
 
func main(){
	fmt.Printf("%+v\n",[]interface{}{partitionDisjoint([]int{6,0,8,30,37,6,75,98,39,90,63,74,52,92,64})})
}