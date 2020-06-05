package main

import "fmt"

func main() {
	n := longestValidParentheses("(()))())(")
	fmt.Printf("%+v\n", []interface{}{n})
}

func longestValidParentheses(s string) int {
	sLen := len(s)
	dp := make([]int, sLen)
	max := 0
	for i := 0; i < sLen; i++ {
		if s[i] == 40 {
			continue
		}
		if i > 0 {
			for j := i; j > 0; {
				j--
				if dp[j] > 0 {
					j = j - dp[j] + 1
				} else if dp[j] == 0 && s[j] == 40 {

					if j > 0 {
						dp[i] = i - j + 1 + dp[j-1]
					} else {
						dp[i] = i - j + 1
					}
					if max < dp[i] {
						max = dp[i]
					}
					break
				}
			}
		}
		fmt.Printf("%+v\n", []interface{}{dp})
	}
	return max
}
