package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {

	if len(input) < 26 {
		return false
	}

	table := make([]bool, 26)

	input = strings.ToLower(input)

	for _, r := range input {
		if unicode.IsLetter(r) {
			table[r-'a'] = true
		}
	}

	for _, v := range table {
		if !v {
			return false
		}
	}

	return true
}

/**
map will cost more time

if len(input) < 26 {
		return false
}

table = make(map[rune]bool)

for _, r := range input {
		if unicode.IsLetter(r) {
			table[r - 'a'] = true
		}
}

return len(table) == 26
BenchmarkPangram-8        100000             14856 ns/op            4140 B/op         49 allocs/op

---------------------------------------------------------------------------------------------------
use array
BenchmarkPangram-8       1000000              1225 ns/op              96 B/op          2 allocs/op

*/
