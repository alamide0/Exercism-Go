package cryptosquare
import (
	"math"
	"strings"
)

func normalizer(r rune) (newR rune) {
	r, newR = r|32, -1

	if 'a' <= r && r <= 'z' || '0' <= r && r <= '9' {
		newR = r
	}
	return
}

func Encode(pt string) string {
	pt = strings.Map(normalizer, pt)

	len := len(pt)

	c, r := getCR(len)
	var b strings.Builder
	
	index := 0

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			index = j * c + i
			if index >= len{
				b.WriteByte(' ')
				continue
			}
			b.WriteByte(pt[index])
		}

		if i != c - 1 {
			b.WriteByte(' ')
		}
	}

	return b.String()
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
