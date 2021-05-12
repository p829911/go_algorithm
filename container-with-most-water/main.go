package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	retVal := 0
	for i < j {
		if height[i] > height[j] {
			if retVal < height[j]*(j-i) {
				retVal = height[j] * (j - i)
			}
			j--
		} else {
			if retVal < height[i]*(j-i) {
				retVal = height[i] * (j - i)
			}
			i++
		}
	}
	return retVal
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
