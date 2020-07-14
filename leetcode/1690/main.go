package main

import (
	"fmt"
)

// Tree .. 前缀树
type Tree struct {
	C    byte
	Next map[byte]*Tree
	Fin  bool
}

var mt map[byte]*Tree

func ctree(dict []string) map[byte]*Tree {
	m := map[byte]*Tree{}
	for _, str := range dict {
		bytes := []byte(str)
		tmp := m
		for index, c := range bytes {
			if _, ok := tmp[c]; !ok {
				tmp[c] = &Tree{
					C:    c,
					Next: map[byte]*Tree{},
				}
			}
			if index == len(bytes)-1 {
				tmp[c].Fin = true
			}
			tmp = tmp[c].Next
		}
	}
	return m
}
func respace(dictionary []string, sentence string) int {
	mt = ctree(dictionary)
	tmp := mt
	count, start, flag := 0, 0, 0
	reset := true
	bytes := []byte(sentence)
	for i := 0; i < len(bytes); i++ {
		v := bytes[i]
		if _, ok := tmp[v]; !ok {

			if reset {
				count += i - flag
				i--
			} else {
				i = flag
			}
			flag, start = i, i
			tmp = mt
			reset = true
		} else {
			if tmp[v].Fin {
				// 如果当前节点是单词结尾
				flag = i
			}
			tmp = tmp[v].Next
			reset = false
		}
		if count > len(sentence)*2 {
			break
		}
		fmt.Printf("%+v\n", []interface{}{i, string(v), count})
	}
	if flag != len(bytes)-1 {
		count += len(bytes) - 1 - start
	}
	return count
}
func main() {
	//jesslookedjustliketimherbrother
	// iwiwwwmuiccwwwwwmumwwwmcciwwuiwcicwwuwicuiwciwmiwicwuwwmuimccwucuuim
	fmt.Printf("%+v\n", []interface{}{respace(
		[]string{"wccm", "wiw", "uwwiwcmiiiwmmwicuwu", "mw"},
		"umwww",
	)})
	// i wiw wwmuiccwwwwwmu mw wwmcciwwuiwcicwwuwicuiwciwmiwicwuwwmuimccwucuuim
}
