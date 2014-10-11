package main

import (
	"fmt"
	"math/big"

	b "github.com/shepherdsm/euler/big"
)

// Power digit sum
// projecteuler.net/problem=16
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int64 {
	return b.DigitSum(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil))
}
