package luhn

import (
	"strings"
	"unicode"
)

func Valid(input string) bool {

	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	result := 0

	runes := []rune(input)
	count := 1
	for i := len(runes) - 1; i >= 0; i-- {
		if !unicode.IsDigit(runes[i]) {
			return false
		}
		if count % 2 == 0 {
			t := (int(runes[i] - '0')) * 2
			if t > 9 {
				result += t - 9
			}else{
				result += t
			}
		}else{
			result += int(runes[i] - '0')
		}

		count++
	}
	return result % 10 == 0
}
