package main

import (
	"fmt"
)

// Uni-modal function
func f(x float64) float64 {
	return -(x-3)*(x-3) + 9
}

func ternarySearch(left, right float64, eps float64) float64 {
	for right-left > eps {
		m1 := left + (right-left)/3
		m2 := right - (right-left)/3

		if f(m1) < f(m2) {
			left = m1
		} else {
			right = m2
		}
	}
	return (left + right) / 2
}

func main() {
	left, right := 0.0, 6.0
	eps := 1e-7
	x := ternarySearch(left, right, eps)
	fmt.Printf("Maximum at x ≈ %.7f, f(x) ≈ %.7f\n", x, f(x))
}

