package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var s string = "    -42"

func myAtoi(s string) int {
	trimS := strings.TrimSpace(s)
	var builder strings.Builder
	minNum := math.MinInt32
	maxNum := math.MaxInt32

	if len(trimS) == 0 {
		return 0
	}
	strList := []rune(trimS)
	firstStr := strList[0]

	if firstStr == '-' || firstStr == '+' {
		builder.WriteRune(firstStr)
	} else if unicode.IsDigit(firstStr) {
		builder.WriteRune(firstStr)
	} else {
		return 0
	}

	for _, v := range strList[1:] {
		if unicode.IsDigit(v) {
			builder.WriteRune(v)
		} else {
			break
		}
	}

	finalNum, _ := strconv.Atoi(builder.String())

	if finalNum < minNum {
		return minNum
	} else if finalNum > maxNum {
		return maxNum
	}
	return finalNum
}

func main() {
	num := myAtoi(s)
	fmt.Println(num)
}
