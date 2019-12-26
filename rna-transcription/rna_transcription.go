package strand

import (
	"strings"
)

var pairs = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	if dna == "" {
		return dna
	}

	return strings.Map(func(r rune) rune {
		return pairs[r]
	}, dna)
}
