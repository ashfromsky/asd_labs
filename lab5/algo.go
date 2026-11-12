package main

func shakerSort(arr []int) []int {
	left := 0
	right := len(arr) - 1
	for left < right {
		for j := left; j < right; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		right--

		for k := right; k > left; k-- {
			if arr[k] < arr[k-1] {
				arr[k], arr[k-1] = arr[k-1], arr[k]
			}
		}
		left++
	}
	return arr
}
