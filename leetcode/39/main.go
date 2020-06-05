package main

import (
	"fmt"
)

var m map[int][][]int

func init() {
	m = make(map[int][][]int)
}
func main() {
	out := combinationSum([]int{7, 3, 2}, 18)
	fmt.Printf("\n%+v\n", []interface{}{out})
}

// DP
func combinationSum(candidates []int, target int) [][]int {
	sorted := quickSort(candidates)
	for i := 0; i <= target; i++ {
		combi(sorted, i)
		// fmt.Printf("%d : %+v\n\n", i, []interface{}{m})
	}
	return m[target]
}

func combi(candidates []int, n int) [][]int {
	if _, already := m[n]; already {
		return m[n]
	}
	cLen := len(candidates)
	for i := 0; i < cLen; i++ {
		if candidates[i] > n {
			break
		}
		diff := n - candidates[i]
		if diff == 0 {
			m[n] = append(m[n], []int{n})
		} else if diff == candidates[i] {
			m[n] = append(m[n], []int{diff, diff})
		} else {
			if _, ok := m[diff]; !ok {
				diffCombi := combi(candidates, diff)
				tmp := [][]int{}
				for k := 0; k < len(diffCombi); k++ {
					item := append([]int{candidates[i]}, diffCombi[k]...)
					tmp = append(tmp, item)
				}
				m[n] = append(m[n], tmp...)
			} else {
				diffCombi := m[diff]
				tmp := [][]int{}

				for k := 0; k < len(diffCombi); k++ {
					item := append([]int{candidates[i]}, diffCombi[k]...)
					tmp = append(tmp, item)
				}
				m[n] = append(m[n], tmp...)
			}
		}
	}
	m[n] = uniq(m[n])
	return m[n]
}

func uniq(in [][]int) [][]int {
	u := map[string]bool{}
	out := [][]int{}
	for _, item := range in {
		flag := false
		keys := []byte{}
		sorted := quickSort(item)
		for k, v := range sorted {
			if k < len(sorted)-1 && v == sorted[k+1] {
				flag = true
				break
			}
			keys = append(keys, byte(v))
		}
		if flag {
			continue
		}
		key := string(keys)
		if !u[key] {
			u[key] = true
			out = append(out, item)
		}
	}
	return out
}

func quickSort(candidates []int) []int {
	cLen := len(candidates)
	if cLen == 0 {
		return []int{}
	}
	if cLen == 1 {
		return candidates
	}
	base := candidates[0]
	left := []int{}
	right := []int{}
	for i := 1; i < cLen; i++ {
		if candidates[i] < base {
			left = append(left, candidates[i])
		} else {
			right = append(right, candidates[i])
		}
	}
	return append(quickSort(left), append([]int{base}, quickSort(right)...)...)
}
