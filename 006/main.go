package main

import "fmt"

func sumOfSquares(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		sum += (i * i)
	}

	return sum
}

func squareOfSum(max int) int {
	sum := 0
	for i := 0; i <= max; i++ {
		sum += i
	}

	return sum * sum
}

func main() {
	problem := `The sum of the squares of the first ten natural numbers is,
	12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,
	(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.`
	max := 100
	s1 := sumOfSquares(max)
	s2 := squareOfSum(max)
	answer := s2 - s1
	work := fmt.Sprintf("%d - %d = %d", s2, s1, answer)
	fmt.Printf("%s\n%s\nAnswer: %d\n", problem, work, answer)
}
