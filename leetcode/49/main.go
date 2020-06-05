package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", []interface{}{groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"})})
}
func groupAnagrams(strs []string) [][]string {
	l := len(strs)
	m := map[string][]string{}
	result := [][]string{}
	for i := 0; i < l; i++ {
		c := cal(strs[i])
		if _, ok := m[c]; !ok {
			m[c] = []string{}
		}
		m[c] = append(m[c], strs[i])
	}
	for _, item := range m {
		result = append(result, item)
	}
	return result
}

func cal(s string) string {
	r := []byte{}
	for i := 0; i < len(s); i++ {
		r = append(r, s[i])
	}
	r = quickSort(r)
	return string(r)
}

func quickSort(candidates []byte) []byte {
	cLen := len(candidates)
	if cLen == 0 {
		return []byte{}
	}
	if cLen == 1 {
		return candidates
	}
	base := candidates[0]
	left := []byte{}
	right := []byte{}
	for i := 1; i < cLen; i++ {
		if candidates[i] < base {
			left = append(left, candidates[i])
		} else {
			right = append(right, candidates[i])
		}
	}
	return append(quickSort(left), append([]byte{base}, quickSort(right)...)...)
}
