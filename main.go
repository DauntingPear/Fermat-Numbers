package main

import (
	"fmt"
	"strconv"
    "math"
)

func isPrime(n int) bool {
    if n > 1 && n <= 3 {
        return true
    }
    if n <= 1 || n % 2 == 0 || n % 3 == 0 {
        return false
    }

    for i := 5; i*i <= n; i += 6 {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }
    }
    return true
}

func main() {
    var input string
    fmt.Print("Enter a number: ")
    fmt.Scanln(&input)
    number, _ := strconv.Atoi(input)

    for i := number; i > 0; i-- {
        fermat := math.Pow(2, math.Pow(2, float64(i))) + 1
        fmt.Print("Fermat number: ", fermat, " -- Is Prime: ", isPrime(int(fermat)), "\n")
    }
}
