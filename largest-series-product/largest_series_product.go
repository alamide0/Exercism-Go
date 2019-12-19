package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}

	for _, r := range digits {
		if r < '0' || r > '9' {
			return -1, errors.New("digits input must only contain digits")
		}
	}

	var max, current int64 = -1, 1

	for i := 0; i <= len(digits)-span; i++ {
		for j := i; j < i+span; j++ {
			current *= int64(digits[j] - '0')
		}
		if current > max {
			max = current
		}
		current = 1
	}

	return max, nil
}

/**
return error with fmt.Errorf()
//BenchmarkLargestSeriesProduct-8          1000000              1103 ns/op             240 B/op          8 allocs/op
return error with errors.New()
//BenchmarkLargestSeriesProduct-8          2000000               676 ns/op              64 B/op          4 allocs/op

so，we sould use errors.New(), other than fmt.Errorf()
*/
