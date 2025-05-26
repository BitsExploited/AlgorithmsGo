/*
Time Complexity: O(n)

(From right) → 7, 6, 4, 6, 3, 4, 3, 0, 8, 8, 4, 1, 9, 3, 5, 4
Double these:    ↑   ↑   ↑   ↑   ↑   ↑   ↑   ↑
Becomes:         7, 12, 4, 12, 3, 8, 3, 0, 8, 16, 4, 2, 9, 6, 5, 8
Then subtract 9 from values > 9:
                 7, 3, 4, 3, 3, 8, 3, 0, 8, 7, 4, 2, 9, 6, 5, 8
Sum: **70**
→ `70 % 10 == 0` → Valid
*/

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
