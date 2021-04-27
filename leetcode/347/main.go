package main

import (
	"fmt"
	"sort"
)

type cp struct {
	Member int
	Score  int
}

type cps []cp

func (c cps) Len() int {
	return len(c)
}

func (c cps) Less(i, j int) bool {
	return c[i].Score > c[j].Score
}

func (c cps) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func topKFrequent(nums []int, k int) []int {
	c := cps{}
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		c = append(c, cp{
			Member: v,
			Score:  m[v],
		})
	}

	sort.Sort(c)
	// fmt.Printf("%+v\n", []interface{}{c})
	result := make([]int, k)
	index := 0
	l := len(c)
	for i := 0; i < l; i++ {
		if i == 0 || result[i-1] != c[i].Member {
			// result = append(result, c[i].Member)
			result[index] = c[i].Member
			index++
			if index >= k {
				return result
			}
		}
	}
	return result
}

func main() {
	// fmt.Printf("%+v\n", []interface{}{topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)})
	fmt.Printf("%+v\n", []interface{}{topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)})
}
