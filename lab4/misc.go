package main

import "fmt"

func findInMatrix(matrix [][]float64) (int, int, bool) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1, -1, false
	}

	cols := len(matrix[0])
	for col := range cols {
		row := binarySearchCol(matrix, col)
		if row != -1 {
			return row, col, true
		}
	}

	return -1, -1, false
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%6.1f ", val)
		}
		fmt.Println()
	}
}
