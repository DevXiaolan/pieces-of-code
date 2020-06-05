package main

import (
	"fmt"
)

const MAX, MIN = int(^uint32(0) >> 1), ^int(^uint32(0) >> 1)

func main() {
	fmt.Printf("%+v\n", myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

func myAtoi(str string) int {
	hasSign := false
	sign := int64(1)
	m := map[string]int64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	nums := []int64{}
	for _, c := range str {
		if string(c) == " " && hasSign == false {
			continue
		} else if string(c) == "-" && hasSign == false {
			sign = -1
			hasSign = true
		} else if string(c) == "+" && hasSign == false {
			sign = 1
			hasSign = true
		} else if val, isPresent := m[string(c)]; isPresent {
			nums = append(nums, val)
			hasSign = true
		} else {
			break
		}

	}

	sum := int64(0)
	for index, bit := range nums {
		sum += bit * pow10(len(nums)-1-index)

		if sum < 0 {
			break
		}
	}

	if sum < 0 || sum > int64(MAX) {
		if sign > 0 {
			return MAX
		}
		return MIN
	}
	return int(sum * sign)
}

func pow10(n int) int64 {
	if n == 0 {
		return 1
	} else {
		r := 10 * pow10(n-1)
		if r > int64(MAX) {
			return int64(MAX)
		}
		return r
	}
}
