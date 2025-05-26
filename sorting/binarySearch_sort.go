package main

import "fmt"

func binarySearchSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		low := 0
		high := i - 1

		for low <= high {
			mid := low + (high-low)/2
			if arr[mid] > temp {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		for j := i; j > low; j-- {
			arr[j] = arr[j-1]
		}

		arr[low] = temp
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	binarySearchSort(arr)

	fmt.Println(arr)
}
