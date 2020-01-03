package atbash

import (
	"strings"
)

func Atbash(s string) string {
	s = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		r |= 32
		if r >= 'a' && r <= 'z' {
			return 'a' + 'z' - r
		}
		return -1
	}, s)

	var builder strings.Builder

	for i := 0; i < len(s); i += 5 {
		if i+5 < len(s) {
			builder.WriteString(s[i : i+5])
			builder.WriteString(" ")
		} else {
			builder.WriteString(s[i:])
		}
	}

	return builder.String()
}
