package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Count int
	Next  map[string]*Node
}

func subdomainVisits(cpdomains []string) []string {
	m := map[string]*Node{}
	for _, v := range cpdomains {
		splits := strings.Split(v, " ")
		c, _ := strconv.Atoi(splits[0])
		domain := splits[1]
		dots := strings.Split(domain, ".")
		t
		for i := len(dots) - 1; i >= 0; i-- {
			if _, ok := m[dots[i]]; !ok {
				m[dots[i]] = &Node{}
			}
			m[dots[i]].Count += c
		}
	}
	return []string{}
}
func main() {
	fmt.Printf("%+v\n", []interface{}{subdomainVisits([]string{"9001 discuss.leetcode.com"})})
}
