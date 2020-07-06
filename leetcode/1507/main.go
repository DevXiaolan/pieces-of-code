package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				break
			}
		}
	}
	return false
}

func main() {

}
