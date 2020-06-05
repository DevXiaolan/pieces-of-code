package main

import (
	"unsafe"
	"fmt"
)

func isPalind(s string) bool {
	l := len(s)
	if l <= 1 {
		return true
	}
	for i := 0; i <= l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

func shortestPalindrome(s string) string {
	l := len(s)
	tmp := 0
	
	for i := 0; i < l; i++ {
		if isPalind(s[0:i+1]) {
			tmp = i
		}
	}
	result := []byte{}
	for i:=l-1;i>tmp;i--{
		result = append(result,s[i])
	}
	result = append(result, s...)

	return *(*string)(unsafe.Pointer(&result))
}

func main() {
	str := "abcd"
	// shortestPalindrome(str)
	fmt.Printf("%+v\n", []interface{}{shortestPalindrome(str)})
}
