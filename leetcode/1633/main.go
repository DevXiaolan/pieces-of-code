package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	c := byte(0)
	l := len(s)
	i := l-1
	for i>=0 {
		fmt.Printf("%+v\n",[]interface{}{c,s[i],c^s[i]})
		c = c^s[i]
		
		i--
		
	}
	return c
}

func main() {
	str := "abaccdeff"
	fmt.Printf("%+v\n", []interface{}{string(firstUniqChar(str))})
}
