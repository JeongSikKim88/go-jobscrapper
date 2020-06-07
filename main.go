package main

import (
	"fmt"
)

func superadd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := superadd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
