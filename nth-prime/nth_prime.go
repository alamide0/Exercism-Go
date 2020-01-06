package prime

import (
	"math"
)

func Nth(n int) (int, bool) {

	if n <= 0 {
		return 0, false
	}

	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i, true
			}
		}
	}

	return 0, false
}

func isPrime(n int) bool {

	if n % 2 == 0 {
		return n == 2
	}

	for i, end := 3, intSqrt(n); i <= end; i+=2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func intSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}
