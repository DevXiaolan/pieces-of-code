package main

import "fmt"

// ( 40 )41 [ 91 ] 93  { 123 } 125
func main() {
	fmt.Printf("%+v\n", []interface{}{isValid("()[]{}")})
}

func isValid(s string) bool {
	fixMap := map[byte]byte{
		40:  41,
		91:  93,
		123: 125,
	}
	stack := []byte{}
	for _, b := range []byte(s) {
		if b == 40 || b == 91 || b == 123 {
			stack = append(stack, b)
		} else if b == 41 || b == 93 || b == 125 {
			if len(stack) == 0 {
				return false
			}
			if fixMap[stack[len(stack)-1]] == b {
				// hit!
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
