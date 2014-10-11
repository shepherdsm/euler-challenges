package main

import (
	"fmt"
	"math/big"
)

// Power digit sum
// projecteuler.net/problem=16
func main() {
	fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() int64 {
	sum := big.NewInt(0)
	rem := big.NewInt(0)
	num := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil)

	for num.Cmp(big.NewInt(0)) > 0 {
		num.DivMod(num, big.NewInt(10), rem)
		sum.Add(sum, rem)
	}

	return sum.Int64()
}
