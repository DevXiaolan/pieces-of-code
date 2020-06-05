package main

import (
	"fmt"
	"strings"
)

var m map[int]string

func main() {
	n := countAndSay(5)
	fmt.Printf("%+v\n", []interface{}{n})
}
func run(n string) string {
	out := []string{}
	nLen := len(n)
	flag := n[0]
	count := 1
	for i := 1; i < nLen; i++ {
		if n[i] == flag {
			count++
		} else {
			out = append(out, fmt.Sprintf("%d%c", count, flag))
			flag = n[i]
			count = 1
		}
	}
	out = append(out, fmt.Sprintf("%d%c", count, flag))
	return strings.Join(out, "")
}
func countAndSay(n int) string {
	out := "1"
	for i := 1; i < n; i++ {
		out = run(out)
		// fmt.Printf("%+v\n", []interface{}{out})
	}
	return out
}
