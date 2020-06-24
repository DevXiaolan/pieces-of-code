package main

import (
	"fmt"
)

func destCity(paths [][]string) string {
	m := map[string]string{}
	for _, item := range paths {
		m[item[0]] = item[1]
	}
	for _, v := range m {
		if _, ok := m[v]; !ok {
			return v
		}
	}
	return ""
}

func main() {
	fmt.Printf("%+v\n", []interface{}{destCity([][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	})})
}
