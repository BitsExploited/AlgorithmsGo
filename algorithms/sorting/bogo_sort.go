/* 
This algorithm sorts the elements in the array randomly.
It has very high time complexity (cause obviously)
*/

package main

import (
	"fmt"
	"math/rand"
)

func isSorted(arr []int) bool {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i + 1] {
			return false
		}
	}
	return true
}

func shuffle(arr []int) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func bogoSort(arr []int, n int) []int {
	for !isSorted(arr) {
		shuffle(arr)
	}

	return arr
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	n := len(arr)

	bogoSort(arr, n)
	
	fmt.Println(arr)
}
