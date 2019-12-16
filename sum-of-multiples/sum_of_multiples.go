package summultiples

func SumMultiples(limit int, divisors ...int) (res int) {
	for i := 1; i < limit; i++ {
		for _, v := range divisors {
			
			if v == 0 {
				continue
			}

			if i < v {
				continue
			}

			if i % v == 0 {
				res += i
				break
			}
		}
	}
	return
}
