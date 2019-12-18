package cryptosquare
import (
	"math"
	"unicode"
)

func Encode(pt string) string {
	var normalized []rune

	for _, r := range pt {
		if unicode.IsLetter(r) {
			normalized = append(normalized, unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			normalized = append(normalized, r)
		}
	}

	len := len(normalized)

	c, r := getCR(len)
	var crypto []rune
	
	index := 0

	for i := 0; i < c; i++ {
		
		for j := 0; j < r; j++ {
			index = j * c + i
			if index >= len{
				crypto = append(crypto, ' ')
				continue
			}

			crypto = append(crypto, normalized[index])
		}

		if i != c - 1 {
			crypto = append(crypto, ' ')
		}
	}

	return string(crypto)
}

func intSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func getCR(n int) (c, r int){
	c = intSqrt(n)
	r = c
	if c * c != n {
		c += 1
	}
	if r * c < n {
		r += 1
	}
	return
}
