package queenattack

import (
	"fmt"
	// "regexp"
)

func CanQueenAttack(w string, b string) (bool, error) {
	// reg, err := regexp.Compile("[a-h][1-8]")
	// if err != nil {
	// 	return false, fmt.Errorf("regex expression error")
	// }

	// if !reg.MatchString(w) || !reg.MatchString(b) {
	// 	return false, fmt.Errorf("invaild position")
	// }

	if len(w) != 2 || len(b) != 2 {
		return false, fmt.Errorf("invaild position")
	}

	if !(isInBorder(w[0]) && isInBorder(w[1]) && isInBorder(b[0]) && isInBorder(b[1])) {
		return false, fmt.Errorf("invaild position")
	}

	if w == b {
		return false, fmt.Errorf("same quare")
	}

	if w[0] == b[0] || w[1] == b[1] {
		return true, nil
	}

	if byteSubAbs(w[0], b[0]) == byteSubAbs(w[1], b[1]) {
		return true, nil
	}

	return false, nil

}

func isInBorder(b byte) bool {
	return b >= 'a' && b <= 'h' || b >= '1' && b <= '8'
}

//byte is uint8,
//var a, b byte = 1, 2
//c := a - b
//c = 255
//byte is unsigned
func byteSubAbs(a byte, b byte) int {
	if a > b {
		return int(a - b)
	}

	return int(b - a)
}

//use regexp
//BenchmarkCanQueenAttack-8         300000              4068 ns/op            2935 B/op         39 allocs/op

//no regexp
//BenchmarkCanQueenAttack-8       50000000                32.2 ns/op             0 B/op          0 allocs/op

//regexp spend more time and space
