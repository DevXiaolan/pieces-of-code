package main

import (
	"strings"
	"fmt"
)

func isCap(c byte) bool {
	if c >= 65 && c <= 96 {
		return true
	}
	return false
}

func detectCapitalUse(word string) bool {
	
	// flag:   0 全小写  1 首字母大写  2全大写
	flag := 0
	l := len(word)
	for i := 0; i < l; i++ {
		if i == 0 && isCap(word[i]) {
			flag = 1
			continue
		}
		if flag == 0 && isCap(word[i]) {
			return false
		}
		if flag == 2 && !isCap(word[i]) {
			return false
		}
		if flag == 1 {
			if isCap(word[i]) {
				if  isCap(word[i-1]) {
					flag = 2
					continue
				}else{
					return false
				}
			} 
		}
	}
	return true
}

func main() {
	str := "usa"
	fmt.Printf("%+v\n", []interface{}{detectCapitalUse(str)})
}
