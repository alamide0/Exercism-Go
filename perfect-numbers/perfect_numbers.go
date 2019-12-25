package perfect

import (
	"errors"
	"math"
)

type Classification uint8

const (
	ClassificationDeficient = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("OnlyPositive")

func Classify(n int64) (res Classification, err error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	if n == 1 {
		return ClassificationDeficient, nil
	}

	var sum int64 = 1
	end := intSqrt(n)
	for i := int64(2); i <= end; i++ {
		if n%i == 0 {
			sum += i
			anotherHalf := n / i
			if anotherHalf != i {
				sum += anotherHalf
			}
		}
	}

	switch {
	case sum == n:
		res = ClassificationPerfect
	case sum > n:
		res = ClassificationAbundant
	case sum < n:
		res = ClassificationDeficient
	}

	return res, nil
}

func intSqrt(n int64) int64{
	return int64(math.Sqrt(float64(n)))
}

