package main

import (
	"fmt"
)

func buddyStrings(A string, B string) bool {
	l := len(A)
	if l != len(B) {
		return false
	}

	diff := []int{}
	m := map[byte]bool{}
	hasDump := false
	for i := 0; i < l; i++ {
		if m[A[i]] {
			hasDump = true
		}
		m[A[i]] = true
		if A[i] != B[i] {
			diff = append(diff, i)
			if len(diff) >2 {
				// 超过2个位置不同，
				return false
			}
		}
	}
	ldiff := len(diff)
	if ldiff ==1 {
		return false
	}
	if ldiff==2 && (A[diff[0]] == B[diff[1]]) && (A[diff[1]] == B[diff[0]]) {
		return true
	}
	if ldiff==0 && hasDump {
		// 两个字符串相同,并且有重复字符
		return true
	}
	return false
}

func main() {
	fmt.Printf("%+v\n", []interface{}{buddyStrings("ab", "ab")})
}
