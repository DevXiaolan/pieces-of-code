package main

import (
	"fmt"
)

var m map[int]int
var bigNumber = 999999

func getP(n int) int {
	if v, ok := m[n]; ok {
		return v
	}
	p := 0
	if n == 1 {
		p = 0
	} else if n%2 == 0 {
		p = 1 + getP(n/2)
	} else {
		p = 1 + getP(3*n+1)
	}
	m[n] = p
	return p
}

func insert(r [][]int, item []int, cap int) [][]int {
	l := len(r)
	for i := 0; i < l; i++ {
		if item[0] < r[i][0] {
			return append(r[:i], append([][]int{item}, r[i:]...)...)[:cap]
		} else if item[0] == r[i][0] && item[1] < r[i][1] {
			return append(r[:i], append([][]int{item}, r[i:]...)...)[:cap]
		}
	}
	return r
}

func getKth(lo int, hi int, k int) int {
	m = map[int]int{}

	r := [][]int{}
	for i:=0;i<k;i++ {
		r = append(r,[]int{bigNumber,bigNumber})
	}
	for i := lo; i <= hi; i++ {
		p := getP(i)
		r = insert(r, []int{p, i}, k)
	}
	
	return r[k-1][1]
}

func main() {
	fmt.Printf("%+v\n", []interface{}{getKth(7, 11, 4)})
}
