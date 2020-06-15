package main

import (
	"fmt"
	"strings"
)

var m map[byte]string

func getValidT9Words(num string, words []string) []string {
	m = map[byte]string{
		'a': "2",
		'b': "2",
		'c': "2",
		'd': "3",
		'e': "3",
		'f': "3",
		'g': "4",
		'h': "4",
		'i': "4",
		'j': "5",
		'k': "5",
		'l': "5",
		'm': "6",
		'n': "6",
		'o': "6",
		'p': "7",
		'q': "7",
		'r': "7",
		's': "7",
		't': "8",
		'u': "8",
		'v': "8",
		'w': "9",
		'x': "9",
		'y': "9",
		'z': "9",
	}
	result := []string{}
	for _, word := range words {
		bytes := []string{}
		l := len(word)
		for i := 0; i < l; i++ {
			bytes = append(bytes, m[word[i]])
		}
		
		if num == strings.Join(bytes, "") {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	fmt.Printf("%+v\n", []interface{}{getValidT9Words("8733", []string{"tree", "used"})})
}
