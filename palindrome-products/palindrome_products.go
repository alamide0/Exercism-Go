package palindrome

import (
	"errors"
	// "strconv"
)

type Product struct {
	palindrome     int
	Factorizations [][2]int
}

func Products(fmin int, fmax int) (Product, Product, error) {

	if fmax < fmin {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	min := Product{
		fmax * fmax,
		[][2]int{},
	}

	max := Product{
		fmin * fmin,
		[][2]int{},
	}

	if min.palindrome < max.palindrome {
		min.palindrome, max.palindrome = max.palindrome, min.palindrome
	}

	hasMirror := false

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			m := i * j
			if (m >= max.palindrome || m <= min.palindrome) && isMirror(m) {
				hasMirror = true

				if m < min.palindrome {
					min.palindrome = m
					min.Factorizations = [][2]int{{i, j}}
				} else if m == min.palindrome {
					min.Factorizations = append(min.Factorizations, [2]int{i, j})
				}

				if m > max.palindrome {
					max.palindrome = m
					max.Factorizations = [][2]int{{i, j}}
				} else if m == max.palindrome {
					max.Factorizations = append(max.Factorizations, [2]int{i, j})
				}

			}
		}
	}

	if !hasMirror {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return min, max, nil
}

func isMirror(p int) bool {
	if p > -10 && p < 10 {
		return true
	}

	revp := 0
	for n := p; n > 0; n /= 10 {
		revp = revp*10 + n%10
	}
	return p == revp

}

// func isMirror(n int) bool {
// 	if n < 10 {
// 		return true
// 	}

// 	s := strconv.Itoa(n)
// 	len := len(s)
// 	end := len / 2
// 	for i := 0; i < end; i++ {
// 		if s[i] != s[len-i-1] {
// 			return false
// 		}
// 	}

// 	return true
// }
