package hamming
import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("the length is not same!")
	}

	if a == b {
		return 0, nil
	}

	diff := 0

	aRunes := []rune(a)
	bRunes := []rune(b)

	for i, r := range aRunes {
		if r != bRunes[i] {
			diff++
		}
	}

	return diff, nil
}
