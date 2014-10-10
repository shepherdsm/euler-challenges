package main

import "fmt"

// Smallest multiple
// projecteuler.net/problem=5
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int {
	return 2 * 3 * 2 * 5 * 7 * 2 * 3 * 11 * 13 * 2 * 17 * 19
}
