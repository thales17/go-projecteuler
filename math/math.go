package math

import (
	"fmt"
	"math"
	"strings"
)

// BiggestPrimeFactor returns the largest prime factor of a number
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

// IsPalindrome returns true if the number is a palindrome
func IsPalindrome(n int) bool {
	nStr := fmt.Sprintf("%d", n)
	nArr := strings.Split(nStr, "")
	for i := 0; i < len(nArr)/2; i++ {
		nArr[i], nArr[len(nArr)-i-1] = nArr[len(nArr)-i-1], nArr[i]
	}
	nRevStr := strings.Join(nArr, "")
	return nStr == nRevStr
}

// Factorial returns the factorial n
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// IsPrime returns true if the number is Prime
func IsPrime(n int) bool {
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
