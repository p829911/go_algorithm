package main

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	diff := strconv.FormatInt(int64(x^y), 2)
	sum := 0
	for _, bit := range diff {
		bitInt, _ := strconv.Atoi(string(bit))
		sum += bitInt
	}
	return sum
}

func main() {
	x, y := 1, 4
	fmt.Println(hammingDistance(x, y))
}
