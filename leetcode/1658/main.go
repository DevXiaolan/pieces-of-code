package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	i, j := 0, 1
	l := len(s)
	m := map[byte]bool{}
	for i < l {
		for j < l {
			if s[i] == s[j] {
				break
			}
			j++
		}
		if j == l && !m[s[i]] {
			return s[i]
		}
		m[s[i]] = true
		i++
		j = i + 1
	}
	return ' '
}

func main() {
	fmt.Printf("%+v\n", []interface{}{string(firstUniqChar("dddccdbba"))})
}
