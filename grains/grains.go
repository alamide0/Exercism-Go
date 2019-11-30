package grains
import "errors"

func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("input is illegal")
	}

	return 1 << uint((input - 1)), nil

}

func Total() uint64 {
	n, _ := Square(64) 

	return n * 2 - 1
}