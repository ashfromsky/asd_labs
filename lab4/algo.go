package main

func binarySearchCol(matrix [][]float64, col int) int {
	left := 0
	right := len(matrix) - 1

	for left <= right {
		middle := (left + right) / 2
		val := matrix[middle][col]

		if val >= 0 && val <= 5 {
			return middle
		} else if val < 0 {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
