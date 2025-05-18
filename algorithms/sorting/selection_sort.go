package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	selectionSort(arr)

	fmt.Println(arr)
}
