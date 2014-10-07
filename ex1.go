package main

import "fmt"

// Multiples of 3 and 5
// projecteuler.net/problem=1
func main() {
    sum := 0

    for i := 1; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }

    fmt.Printf("The answer is %d\n", sum)
}
