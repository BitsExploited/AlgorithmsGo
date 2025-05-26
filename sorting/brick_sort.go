package main

import "fmt"

func brickSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	swap := true
	for swap {
		swap = false

		//Odd-indexed pairs
		for i := 1; i < len(arr) - 1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}

		// Even-indexed pairs
		for i := 0; i < len(arr) - 1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	brickSort(arr)

	fmt.Println(arr)
}
