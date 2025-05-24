package main

import "fmt"

func interpolationSearch(arr []int, low, high, x int) int {
	if low <= high && x >= arr[low] && x <= arr[high] {
		pos := low + ((high-low)/(arr[high]-arr[low]))*(x-arr[low])
		if arr[pos] == x {
			return pos
		}
		if arr[pos] < x {
			return interpolationSearch(arr, pos+1, high, x)
		}
		if arr[pos] > x {
			return interpolationSearch(arr, low, pos-1, x)
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 7, 8, 11, 12}
	n := len(arr)
	x := 12
	index := interpolationSearch(arr, 0, n-1, x)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", x, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", x)
	}
}
