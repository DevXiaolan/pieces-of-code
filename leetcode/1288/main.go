package main

import (
	"fmt"
)

func sort(intervals [][]int) {
	l := len(intervals)
	for i:=0;i<l;i++ {
		for j:=i+1;j<l;j++ {
			if intervals[i][0] > intervals[j][0] {
				// swap
				intervals[i],intervals[j] = intervals[j],intervals[i]
			}else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1]{
				// swap
				intervals[i],intervals[j] = intervals[j],intervals[i]
			}
		}
	}
}

func removeCoveredIntervals(intervals [][]int) int {
	sort(intervals)
	fmt.Printf("%+v\n",[]interface{}{intervals})
	prev := 0
	count :=0
	for _,v := range intervals {
		if v[1]<=prev {
			count++
		}else{
			prev = v[1]
		}
	}
	return len(intervals)-count
}

func main() {
	fmt.Printf("%+v\n", []interface{}{removeCoveredIntervals([][]int{
		// []int{1, 2},
		// []int{1,4},
		// []int{3,4},
		[]int{1,4},
		[]int{3,6},
		[]int{2,8},
	})})
}
