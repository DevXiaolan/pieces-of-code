package main

import "fmt"

func replaceSpace(s string) string {
	bytes := []byte(s)
	res := []byte{}
	for _, v := range bytes {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
func main() {
	fmt.Printf("%+v\n", []interface{}{replaceSpace("We are happy.")})
}
