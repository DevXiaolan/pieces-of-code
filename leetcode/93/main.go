package main

import (
	"fmt"
)

func find(s string, n int) []string {
	l := len(s)
	if l == n*3 {
		return []string{string(s[0:3]), string(s[3:6]), string(s[6:9]), string(s[9:12])}
	}
}
func restoreIpAddresses(s string) []string {

	return find(s, 4)
}

func main() {
	fmt.Printf("%+v\n", []interface{}{restoreIpAddresses("255255111135")})
}
