package main

import "fmt"

func countingSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	count := make([]int, max + 1)

	for _, num := range arr {
		count[num]++
	}

	i := 0
	for val := 0; val <= max; val++ {
		for count[val] > 0 {
			arr[i] = val
			i++
			count[val]--
		}
	}
}

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	countingSort(arr)

	fmt.Println(arr)
}
