
/*
Time Complexity
Best Case: O(n) //  If optimised with the "swapped" flag
Average Case: O(n^2)
Worst Case: O(n^2)

*/

package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	bubbleSort(arr)

	fmt.Println(arr)

}
