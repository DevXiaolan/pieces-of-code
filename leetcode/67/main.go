package main

import (
	"fmt"
)

func add(a byte, b byte, c byte) (byte, byte) {
	switch a + b + c - 48*3 {
	case 0:
		return '0', '0'
	case 1:
		return '0', '1'
	case 2:
		return '1', '0'
	case 3:
		return '1', '1'
	}
	return '0', '0'
}
func addBinary(a string, b string) string {
	result := []byte{}
	i, j, sig := len(a)-1, len(b)-1, byte(48)
	for i >= 0 || j >= 0 {
		var n, x, y, z byte = '0', '0', '0', sig
		if i >= 0 {
			x = a[i]
		}
		if j >= 0 {
			y = b[j]
		}
		sig, n = add(x, y, z)
		i--
		j--
		result = append([]byte{n}, result...)
	}
	if sig == 49 {
		result = append([]byte{sig}, result...)
	}
	return string(result)
}

func main() {
	fmt.Printf("%+v\n", []interface{}{addBinary("1010", "1011")})
}
