package main

import "fmt"

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j <= end-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[end], arr[i] = arr[i], arr[end]
	return i
}

func quickSort(arr []int, start int, end int) {
	if end <= start {
		return
	}

	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n := len(arr)

	quickSort(arr, 0, n-1)

	fmt.Println(arr)
}
