/*
Best Case: O(n log n) When the array is partially sorted and a good gap sequence

There are lots of worst cases for this sorting algo:

    Depends on gap sequence.

    For original Shell's gap sequence (n/2, n/4, ..., 1):

        Worst Case: O(n²)

    For Hibbard’s sequence (1, 3, 7, 15...):

        Worst Case: O(n^(3/2))

    For Sedgewick’s sequence:

        Worst Case: O(n^(4/3))

    For Tokuda's sequence:

        Worst Case: ~O(n^(1.2))
*/
package main

import "fmt"

func shellSort(arr []int) {
	n := len(arr)

	for gap := n / 2; gap > 0; gap /= 2 {
		// Performing insertion sort for this gap sequence
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j - gap] > temp {
				arr[j] = arr[j - gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	shellSort(arr)

	fmt.Println(arr)
}
