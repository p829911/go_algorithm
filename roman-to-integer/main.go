package main

import "fmt"

func romanToInt(s string) int {
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	sRunes := []rune(s)

	for i, s := range sRunes {
		if i < len(sRunes)-1 {
			if roman[s] < roman[sRunes[i+1]] {
				res -= roman[s]
			} else {
				res += roman[s]
			}
		} else {
			res += roman[s]
		}
	}
	return res
}

func main() {
	stringList := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for _, s := range stringList {
		fmt.Println(romanToInt(s))
	}
}
