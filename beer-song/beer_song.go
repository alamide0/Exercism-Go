package beer

import (
	"errors"
	"fmt"
)

func Verse(verse int) (res string, err error) {
	if verse < 0 || verse > 99 {
		return res, errors.New("illegal verse")
	}

	return format(verse), nil
}

func Verses(upperBound, lowerBound int) (res string, err error) {

	if lowerBound < 0 || upperBound > 99 || lowerBound > upperBound {
		return res, errors.New("illegal bound")
	}

	for i := upperBound; i >= lowerBound; i-- {
		t, err := Verse(i)
		if err != nil {
			return "", err
		}
		res += t + "\n"
	}

	return res, nil
}

func format(n int) string {
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"
	}

	if n == 1 {
		return fmt.Sprintf("1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n")
	}

	if n == 2 {
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", n, n, n-1)
	}

	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
}
