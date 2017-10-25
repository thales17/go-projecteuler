package main

import (
	"fmt"

	"github.com/thales17/go-projecteuler/math"
)

func main() {
	problem := `The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.`
	n := 2000000
	answer := 0
	for i := 2; i < n; i++ {
		if math.IsPrime(i) {
			answer += i
		}
	}

	fmt.Printf("%s\nAnswer: %d\n", problem, answer)
}
