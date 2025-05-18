package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (low + high) / 2

		if arr[mid] == target {
			return mid // Found the element
		} else if arr[mid] < target {
			low = mid + 1 // Search the right side
		} else if arr[mid] > target {
			high = mid - 1 // Search the left side
		}
	}
	return -1 // Not found
}

func bubbleSort(arr []int, n int) {
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{12, 43, 56, 2, 34, 1, 67, 43}
	n := len(arr)
	//Sorting because binary sort works only on a sorted list
	bubbleSort(arr, n)
	
	target := 34

	index := binarySearch(arr, target)

	if index == -1 {
		fmt.Printf("The element %v is not found", target)
	} else {
		fmt.Printf("The element %v is found at the index %v", target, index)
	}
}
