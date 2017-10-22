package main

import "fmt"

func bruteForce(maxDivisible int, start int) int {
	answer := -1
	for i := start; i > 0; i -= maxDivisible {
		pass := true
		for j := 1; j <= maxDivisible; j++ {
			if i%j != 0 {
				pass = false
				break
			}
		}
		if pass {
			answer = i
			fmt.Printf("Answer candidate: %d\n", answer)
		}
	}

	return answer
}

func altMethod(maxDivisible int) int {
	answer := 2
	currentDivisible := 2
	for currentDivisible <= maxDivisible {
		if answer < currentDivisible {
			answer *= currentDivisible
		} else if answer%currentDivisible != 0 {
			answer *= currentDivisible
			fmt.Println(answer)
		}
		currentDivisible++
	}

	return bruteForce(maxDivisible, answer)
}

func main() {
	problem := `2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?`

	maxDivisible := 20
	fmt.Printf("%s\nAnswer:%d\n", problem, altMethod(maxDivisible))
}
