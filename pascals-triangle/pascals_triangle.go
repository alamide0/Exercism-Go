package pascal

var cache = [][]int {
	{1},
	{1, 1},
}

func Triangle(n int)[][]int {

	if n <= len(cache) {//reuse cache
		return cache[:n]
	}

	for i := len(cache); i < n; i++ {//reuse cache
		cache = append(cache, nil)//allocate memory

		cache[i] = append(cache[i], 1)//head

		for j := 0; j < len(cache[i - 1]) - 1; j++ {// add 
			cache[i] = append(cache[i], (cache[i - 1][j] + cache[i - 1][j + 1]))
		}

		cache[i] = append(cache[i], 1)//tail
	}

	return cache
}
//use cache
//BenchmarkPascalsTriangleFixed-8         1000000000               2.21 ns/op            0 B/op          0 allocs/op
//BenchmarkPascalsTriangleIncreasing-8    30000000                49.2 ns/op             0 B/op          0 allocs/op
//no cache
//BenchmarkPascalsTriangleFixed-8           300000              3437 ns/op            5104 B/op         90 allocs/op
//BenchmarkPascalsTriangleIncreasing-8       50000             30543 ns/op           36640 B/op        815 allocs/op