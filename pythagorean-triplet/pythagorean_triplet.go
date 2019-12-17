package pythagorean

import "math"

type Triplet []int

/**
0 < a < sum / 3
a < b < sum / 2

if
	a + b < sum / 2 ----> c > sum / 2 > a + b
because
	a**2 + b**2 < (a + b)**2 < c**2
	c > sum / 3
so
	a + b > sum / 2
	a + b > sum / 3 * 2
---------------------------------------------
according
	a + b + c = sum
	a**2 + b**2 = c**2
get
	sum*(a + b - sum / 2) = a * b
**/
func Sum(sum int) []Triplet {

	var res []Triplet

	d2, d3, tmp := sum/2, sum/3, sum/3*2

	for a := 1; a < d3; a++ {
		for b := a + 1; b < d2; b++ {
			if a + b < d2 {
				continue
			}
			if a + b >= tmp  {
				break
			}

			if sum * (a + b - d2) == a * b {
				res = append(res, Triplet{a, b, sum - a - b})
			}
		}
	}

	return res
}

func Range(min int, max int) []Triplet {
	var res []Triplet
	sum, c := 0, 0
	for a := min; a < max - 1; a++ {
		for b := a + 1; b < max; b++ {
			sum = a * a + b * b
			c = int(math.Sqrt(float64(sum)))

			if c > max {
				break
			}

			if sum == c * c {
				res = append(res, Triplet{a, b, c})
			}
	
		}
	}

	return res
}
