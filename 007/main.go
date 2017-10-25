package main

import (
	"fmt"

	"github.com/thales17/go-projecteuler/math"
)

func findNPrime(n int) int {
	primeCount := 0
	i := 0
	for primeCount <= n {
		i++
		if math.IsPrime(i) {
			fmt.Printf("Prime: %d\n", i)
			primeCount++
		}
	}

	return i
}

func main() {
	problem := `By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
	What is the 10 001st prime number?`
	answer := findNPrime(10001)
	fmt.Printf("%s\nAnswer: %d\n", problem, answer)
}
