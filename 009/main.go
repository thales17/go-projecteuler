package main

import (
	"fmt"
	"math"
)

func main() {
	problem := `A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
	
	a^2 + b^2 = c^2
	For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
	
	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.`

	sum := 1000
	answer := 0 // For problem example it should be 60

	for a := 1; a < sum; a++ {
		pass := false
		for b := a; b < sum; b++ {
			c := math.Sqrt(float64(a*a) + float64(b*b))
			if math.Ceil(c) == c && a+b+int(c) == sum {
				answer = a * b * int(c)
				fmt.Printf("a: %d, b: %d, c: %d\n", a, b, int(c))
				pass = true
			}
		}

		if pass {
			break
		}
	}

	fmt.Printf("%s\nAnswer: %d\n", problem, answer)
}
