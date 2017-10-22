package main

import "fmt"

func isPrime(n int) bool {
	/*
			    if n == 2:
		        return True
		    if n == 3:
		        return True
		    if n % 2 == 0:
		        return False
		    if n % 3 == 0:
		        return False

		    i = 5
		    w = 2

		    while i * i <= n:
		        if n % i == 0:
		            return False

		        i += w
		        w = 6 - w

		    return True
	*/
	if n == 2 {
		return true
	}

	if n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	if n%3 == 0 {
		return false
	}

	i, w := 5, 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}
	return true
}

func findNPrime(n int) int {
	primeCount := 0
	i := 0
	for primeCount <= n {
		i++
		if isPrime(i) {
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
