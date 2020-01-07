package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	fields := strings.Fields(question)
	lastField := fields[len(fields)-1]

	if len(fields) < 3 || fields[0] != "What" || fields[1] != "is" || lastField[len(lastField)-1] != '?' {
		return 0, false
	}
	fields[len(fields)-1] = lastField[:len(lastField)-1]

	res, err := strconv.Atoi(fields[2])

	if err != nil {
		return 0, false
	}

	for i := 3; i < len(fields); {
		n := 0
		switch fields[i] {
		case "plus", "minus":
			if i+1 >= len(fields) {
				return 0, false
			}
			n, err = strconv.Atoi(fields[i+1])
			if fields[i] == "plus" {
				res += n
			} else {
				res -= n
			}
			i += 2
		case "multiplied", "divided":
			if i+2 >= len(fields) {
				return 0, false
			}
			n, err = strconv.Atoi(fields[i+2])
			if fields[i] == "multiplied" {
				res *= n
			} else {
				res /= n
			}
			i += 3
		default:
			return 0, false
		}

		if err != nil {
			return 0, false
		}
	}

	return res, true
}
