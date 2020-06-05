package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(prefix) && j < len(strs[i]); j++ {
			if strs[i][j] != prefix[j] {
				break
			}
		}
		prefix = prefix[0:j]
	}
	return prefix
}
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
