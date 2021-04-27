package main

import "fmt"

func main() {
	m := map[string]bool{}
	p := m

	fmt.Printf("%+v\n", []interface{}{&m, &p})
}
