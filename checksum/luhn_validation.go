package main

import (
	"fmt"
	"unicode"
)

func luhnCheck(number string) bool {
	sum := 0
	double := false

	for i := len(number) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(number[i])) {
			continue
		}
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}

func main() {
	number := "4539148803436467"
	if luhnCheck(number) {
		fmt.Println("Valid Luhn Number")
	} else {
		fmt.Println("Invalid Luhn Number")
	}
}
