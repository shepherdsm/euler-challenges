package main

import (
	"fmt"

	"github.com/shepherdsm/euler/primes"
)

// Summation of primes
// projecteuler.net/problem=10
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() uint {
	var sum uint

	for _, num := range primes.Under(2000000) {
		sum += num
	}

	return sum
}
