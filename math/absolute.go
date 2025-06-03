package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	n1 := 10
	n2 := -10

	ans1 := abs(n1)
	ans2 := abs(n2)

	fmt.Println(ans1)
	fmt.Println(ans2)
}
