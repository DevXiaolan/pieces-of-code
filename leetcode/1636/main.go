package main

import (
	"fmt"
)

func getSum(bits []byte) int {
	sum := 0
	for _, v := range bits {
		vInt := int(v)
		sum += vInt * vInt * vInt * vInt
	}
	return sum
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	resultTmp := map[int][]string{}
	for _, item := range strs {

		key := getSum([]byte(item))
		if _, ok := resultTmp[key]; ok {
			resultTmp[key] = append(resultTmp[key], item)
		} else {
			resultTmp[key] = []string{item}
		}
	}
	for _, v := range resultTmp {
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Printf("%+v\n", []interface{}{groupAnagrams([]string{"run", "had", "lot", "kim", "fat", "net", "fin", "rca", "chi", "lei", "lox", "iva", "liz", "hug", "hot", "irk", "lap", "tan", "tux", "yuk", "hep", "map", "ran", "ell", "kit", "put", "non", "aol", "add", "lad", "she", "job", "mes", "pen", "vic", "fag", "bud", "ken", "nod", "jam", "coy", "hui", "sue", "nap", "ton", "coy", "rut", "fit", "cut", "eta", "our", "oho", "zip"})})
}
