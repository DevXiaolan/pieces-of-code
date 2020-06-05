package main

import "fmt"

func main() {
	r := strStr("mississippi", "pi")
	fmt.Printf("%+v\n", []interface{}{r})
}

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen == 0 {
		return 0
	}
	if haystackLen < needleLen {
		return -1
	}
	if haystack == needle {
		return 0
	}

	for i := 0; i < haystackLen; i++ {
		if haystack[i] != needle[0] {
			continue
		} else {
			if i+needleLen <= haystackLen && needle == haystack[i:i+needleLen] {
				return i
			}
		}
	}
	return -1
}
