package main

import "fmt"

// Multiples of 3 and 5
// projecteuler.net/problem=1
func main() {
    fmt.Printf("Solution 1: %d\n", solve1());
    fmt.Printf("Solution 2: %d\n", solve2());
}

func solve1() (sum int) {
    for i := 1; i <= 999; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }

    return
}

func solve2() (int) {
    natSum := func (target int) (int) {
        n := 999 / target

        return (target * n * (n + 1)) / 2
    }

    return natSum(3) + natSum(5) - natSum(15)
}
