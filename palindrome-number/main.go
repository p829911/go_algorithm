package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	reverse := 0
	for tmp := x; tmp != 0; tmp /= 10 {
		reverse = reverse*10 + tmp%10
	}
	return reverse == x
}

func main() {
	x := []int{121, -121, 10, -101, 11}
	for _, v := range x {
		fmt.Println(isPalindrome(v))
	}
}
