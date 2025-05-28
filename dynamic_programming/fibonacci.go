package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return n
	}

	var n1, n2 int = 0, 1

	for i := 1; i < n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
	}

	return n2
}

func main() {
	n := 10

	ans := fibonacci(n)

	fmt.Println(ans)


}
