package main

import (
	"fmt"
	"strings"
)

type Roman struct {
	Num    int
	Letter string
}

func intToRoman(num int) string {
	romanDict := []Roman{
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
	var result []string

	for _, n := range romanDict {
		for num >= n.Num {
			num -= n.Num
			result = append(result, n.Letter)
		}
	}
	return strings.Join(result, "")
}

func main() {
	numList := []int{3, 4, 9, 58, 1994}
	for _, num := range numList {
		fmt.Println(intToRoman(num))
	}
}
