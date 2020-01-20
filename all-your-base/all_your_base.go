package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (res []int, err error) {

	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	n := 0

	for i, v := range inputDigits {

		if v < 0 || v >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}

		if i == len(inputDigits)-1 {
			n = n + v
		} else {
			n = (n + v) * inputBase
		}
	}

	for {
		res = append(res, n%outputBase)
		n /= outputBase
		if n == 0 {
			break
		}
	}

	for i, end := 0, len(res)/2; i < end; i++ {//reverse
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res, nil
}
