package main

import (
	"fmt"
	"math/rand"
)

const (
	rows = 7
	cols = 8
)

func printMatrix(title string, a [][]int) {
	fmt.Println(title)
	for _, row := range a {
		for _, value := range row {
			fmt.Printf("%4d", value)
		}
		fmt.Println()
	}
}

func isSortedRow(row []int) bool {
	for i := 1; i < len(row); i++ {
		if row[i-1] > row[i] {
			return false
		}
	}
	return true
}

func runTest(title string, a [][]int) {
	fmt.Println(title)
	printMatrix("До:", a)

	for i := range a {
		shakerSort(a[i])
	}

	printMatrix("Після:", a)

	ok := true
	for _, row := range a {
		if !isSortedRow(row) {
			ok = false
			break
		}
	}

	if ok {
		fmt.Println("Результат: OK")
	} else {
		fmt.Println("Результат: ПОМИЛКА")
	}
	fmt.Println()
}

func main() {
	test1 := [][]int{
		{-20, -15, -10, -5, 0, 5, 10, 15},
		{-18, -12, -6, 0, 6, 12, 18, 24},
		{-14, -8, -2, 4, 10, 16, 22, 28},
		{-11, -7, -3, 1, 5, 9, 13, 17},
		{-9, -4, 0, 3, 8, 12, 19, 25},
		{-6, -2, 1, 7, 11, 14, 20, 27},
		{-3, 0, 4, 8, 13, 18, 23, 30},
	}

	rng := rand.New(rand.NewSource(20261113))
	test2 := make([][]int, rows)
	for i := range test2 {
		test2[i] = make([]int, cols)
		for j := range test2[i] {
			test2[i][j] = rng.Intn(101) - 50
		}
	}

	test3 := [][]int{
		{15, 10, 5, 0, -5, -10, -15, -20},
		{24, 18, 12, 6, 0, -6, -12, -18},
		{28, 22, 16, 10, 4, -2, -8, -14},
		{17, 13, 9, 5, 1, -3, -7, -11},
		{25, 19, 12, 8, 3, 0, -4, -9},
		{27, 20, 14, 11, 7, 1, -2, -6},
		{30, 23, 18, 13, 8, 4, 0, -3},
	}

	test4 := [][]int{
		{-20, -10, -5, 0, 7, 12, 18, 30},
		{18, -7, 42, 0, -31, 9, 9, -2},
		{30, 18, 12, 7, 0, -5, -10, -20},
		{4, 4, -2, 4, 0, -2, 4, 0},
		{9, 9, 9, 9, 9, 9, 9, 9},
		{1, -1, 3, 3, 0, -1, 8, 3},
		{42, 0, -42, 17, 17, -8, 5, 5},
	}

	runTest("Тест 1: рядки вже відсортовані", test1)
	runTest("Тест 2: випадкові числа", test2)
	runTest("Тест 3: рядки відсортовані у зворотному порядку", test3)
	runTest("Тест 4: змішані випадки", test4)
}
