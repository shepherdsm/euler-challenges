package main

import "fmt"

// Sum square difference
// projecteuler.net/problem=6
func main() {
    fmt.Printf("Solution 1: %d\n", solve1())
    fmt.Printf("Solution 2: %d\n", solve2())
}

func sumTheSquares (upper int) (sum int) {
    for i := 1; i <= upper; i++ {
        sum += i * i
    }

    return
}

func squareTheSum (upper int) (int) {
    return (upper * upper * (upper + 1) * (upper + 1))/ 4
}

func solve1() (int) {
    upper := 100

    sumSquares := sumTheSquares(upper)
    squareSum := squareTheSum(upper)

    return squareSum - sumSquares
}

func solve2() (int) {
    upper := 100

    sumSquares := (upper * (2 * upper + 1) * (upper + 1)) / 6
    squareSum := squareTheSum(upper)

    return squareSum - sumSquares
}
