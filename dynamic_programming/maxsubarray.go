package main

import (
	"fmt"
)

func maxSubArray(list []int) int {
	maxSum := list[0]
	currentSum := list[0]

	for i := 1; i < len(list); i++ {
		currentSum = max(list[i], currentSum + list[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func main() {
	list := []int{12, 78, 67, 45, 80, 32, 12, 54, 67, 90}

	max := maxSubArray(list)

	fmt.Println(max)
}
