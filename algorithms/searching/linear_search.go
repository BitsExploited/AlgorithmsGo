package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{12, 43, 56, 2, 34, 1, 67, 43}
	target := 34

	index := linearSearch(arr, target)

	fmt.Println("The index of the target:", index)

}
