package main

import (
	"fmt"
	"github.com/shepherdsm/euler/primes"
)

// 10001st prime
// projecteuler.net/problem=7
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() uint {
	cand := uint(3)
	count := 2

	for count < 10001 {
		cand += 2

		if primes.IsPrime(cand) {
			count += 1
		}
	}

	return cand
}
