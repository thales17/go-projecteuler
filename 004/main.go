package main

import (
	"fmt"

	"github.com/thales17/go-projecteuler/math"
)

func main() {
	problem := `A palindromic number reads the same both ways. 
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.`

	answer := 0
	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			p := i * j
			if math.IsPalindrome(p) && p > answer {
				answer = p
			}
		}
	}
	fmt.Printf("%s\nAnswer:%d\n", problem, answer)
}
