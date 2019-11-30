package isogram
import "strings"

func IsIsogram(input string) bool {

	if input == "" {
		return true
	}

	input = strings.ToLower(input)

	for i, r := range input {

		if r == ' ' || r == '-' {
			continue
		}

		if strings.ContainsRune(input[i+1:], r) {
			return false
		}
	}

	return true
}