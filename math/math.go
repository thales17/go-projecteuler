package math

import "math"

// BiggestPrimeFactor finds the largest prime factor of a number
func BiggestPrimeFactor(n int) int {
	// I referenced this algorithm: http://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
	factors := []int{}
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors[len(factors)-1]
}
