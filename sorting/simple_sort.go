package main

import "fmt"

func simpleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr) - 1; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	simpleSort(arr)

	fmt.Println(arr)
}
