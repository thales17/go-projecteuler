package main

import (
	"fmt"

	"github.com/thales17/go-projecteuler/math"
)

func main() {
	problem := `The prime factors of 13195 are 5, 7, 13 and 29.
 What is the largest prime factor of the number 600851475143 ?`
	answer := math.BiggestPrimeFactor(600851475143)
	fmt.Printf("%s\nAnswer:%d\n", problem, answer)
}
