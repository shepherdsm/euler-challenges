package main

import (
	"fmt"
	b "github.com/shepherdsm/euler/big"
	"math/big"
)

// Factorial digit sum
// projecteuler.net/problem=20
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int64 {
	return b.DigitSum(big.NewInt(0).MulRange(2, 100))
}
