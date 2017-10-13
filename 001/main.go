package main

import (
	"fmt"
)

func main() {
	problem := `If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.`
	sum := 0
	for i := 0; i < 334; i++ {
		fmt.Printf("%d..", i*3)
		sum += (i * 3)
		if i < 200 && i%3 != 0 {
			fmt.Printf("%d..", i*5)
			sum += (i * 5)
		}
	}

	fmt.Printf("\n%s\nAnswer: %d\n", problem, sum)
}
