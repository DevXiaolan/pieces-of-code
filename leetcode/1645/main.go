package main

import (
	"fmt"
)

func maxValue(grid [][]int) int {
	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1
	x,y := maxX,maxY
	sum := grid[y][x]
	
	for x>0 || y>0 {
		if x == 0 {
			y--
		} else if y == 0 {
			x--
		} else {
			if grid[y-1][x] > grid[y][x-1] {
				y--
			} else {
				x--
			}
		}
		sum += grid[y][x]
	}
	return sum
}

func main() {
	fmt.Printf("%+v\n", []interface{}{maxValue([][]int{
		[]int{1, 2,5},
		[]int{3,2, 1},
		// []int{4, 2, 1},
	})})
}
