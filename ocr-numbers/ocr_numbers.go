package ocr

import (
	"strconv"
	"strings"
)

const (
	NUM = `
 _     _  _     _  _  _  _  _  
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|
                                
`
)

var nums [10][4][3]byte

func init() {
	strs := strings.Split(NUM, "\n")
	for i := 0; i < 10; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 3; k++ {
				nums[i][j][k] = strs[j+1][i*3+k]
			}
		}
	}
}

func Recognize(in string) []string {
	strs := strings.Split(in, "\n")
	strs = strs[1:]

	numLines := len(strs) / 4

	res := make([]string, numLines)

	for i := 0; i < numLines; i++ {
		start := i * 4
		res[i] = recognizeDigit(strs[start : start+4])
	}

	return res
}

func recognizeDigit(strs []string) (res string) {
	ns, count := len(strs[1])/3, 0
	tns := [4][3]byte{}

	for count < ns {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				tns[i][j] = strs[i][3*count+j]
			}
		}
		res += getNumString(tns)
		count++
	}

	return
}

func getNumString(ns [4][3]byte) string {
	for i := 0; i < 10; i++ {
		if isEqual(nums[i], ns) {
			return strconv.Itoa(i)
		}
	}

	return "?"
}

func isEqual(n1, n2 [4][3]byte) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if n1[i][j] != n2[i][j] {
				return false
			}
		}
	}

	return true
}
