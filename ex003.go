package main

import "fmt"

// Largest prime factor
// projecteuler.net/problem=3
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int {
	num := 600851475143
	start := 3

	for start*start < num {
		if num%start == 0 {
			num /= start
		} else {
			start += 2
		}
	}

	return num
}
