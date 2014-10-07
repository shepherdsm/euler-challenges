package main

import "fmt"

// Sum square difference
// projecteuler.net/problem=6
func main() {
    fmt.Printf("Solution 1: %d\n", solve1())
}

func solve1() (int) {
    upper := 100
    sumSquares := 0

    for i := 1; i <= upper; i++ {
        sumSquares += i * i
    }

    squareSum := (upper * (upper + 1)) / 2

    return squareSum * squareSum - sumSquares
}
