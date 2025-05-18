package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	insertionSort(arr)

	fmt.Println(arr)
}
