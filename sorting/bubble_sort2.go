package main

import "fmt"

func bubbleSort2(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1; i++ {
		swapped := false

		for j := 0; j < n - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	bubbleSort2(arr)

	fmt.Println(arr)
}
