package main

import (
	"errors"
	"fmt"
	"strings"
)

type numeral struct {
	val int
	sym string
}

var nums = []numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func romanToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	var n int
	for _, i := range nums {
		if strings.HasPrefix(str, i.sym) {
			n += i.val
			str = str[len(i.sym):]
		}
	}
	if len(str) > 0 {
		return 0, errors.New("Enter a valid roman numeral")
	}
	return n, nil
}

func main() {
	str := "MDXLVI"

	n, err := romanToInt(str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
}
