package sieve

func Sieve(limit int) (res []int) {
	marked := make([]bool, limit+1)

	n := limit / 2
	for i := 2; i <= n; i++ {
		if marked[i] {
			continue
		}
		for j := 2; ; j++ {
			k := i * j
			if k > limit {
				break
			}
			marked[k] = true
		}
	}

	for i := 2; i <= limit; i++ {
		if !marked[i] {
			res = append(res, i)
		}
	}

	return
}

/**
first version
for j := 2; j <= limit/i; j++ {
			marked[i*j] = true
}

second version
end := limit/i
for j := 2; j <= end; j++ {
			marked[i*j] = true
}

third version
for j := 2; ; j++ {
			k := i * j
			if k > limit {
				break
			}
			marked[k] = true
}

first version
BenchmarkSieve-8          100000             19308 ns/op            5336 B/op         22 allocs/op

second version
BenchmarkSieve-8          300000              5104 ns/op            5336 B/op         22 allocs/op

third version
BenchmarkSieve-8          300000              3610 ns/op            5336 B/op         22 allocs/op

so amazing, save so many time
*** limit/i will cost more time ******
*/
