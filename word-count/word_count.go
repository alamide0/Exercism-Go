package wordcount

import (
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	phrase = strings.ToLower(phrase)

	fields := strings.FieldsFunc(phrase, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '\'')
	})
	res := Frequency{}

	for _, v := range fields {
		if v[0] == '\'' {
			res[string(v[1:len(v)-1])] += 1
		} else {
			res[v] += 1
		}
	}

	return res
}
