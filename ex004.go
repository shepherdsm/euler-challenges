package main

import (
	"fmt"
	"github.com/shepherdsm/euler"
)

// Largest palindrome product
// projecteuler.net/problem=4
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int {
	largest := 0
	for i := 999; i > 300; i-- {
		for k := i; k > 300; k-- {
			tmp := k * i
			if euler.IsPalindrome(tmp) && largest < tmp {
				largest = tmp
			}
		}
	}

	return largest
}
