package main

import (
	"fmt"

	"github.com/shepherdsm/euler/fib"
)

func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() uint {
	var upper uint = 4000000
	var sum uint = 0

	for _, val := range fib.Under(upper) {
		if val%2 == 0 {
			sum += val
		}
	}

	return sum
}
