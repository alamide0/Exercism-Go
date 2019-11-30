// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob
import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {

	isEndWithQ := false
	hasLower := false
	hasCapital := false

	str := strings.TrimSpace(remark)

	if len(str) == 0 {
		return "Fine. Be that way!"
	}

	for _, r := range str {
		isEndWithQ = false
		if isCapital(int(r)) {
			hasCapital = true
			continue
		}else if isLower(int(r)){
			hasLower = true
			continue
		}
		isEndWithQ = r == '?'
	}

	if isEndWithQ {
		if hasCapital && !hasLower{
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}else {
		if hasCapital && !hasLower{
			return "Whoa, chill out!"
		}
	}
	return "Whatever."
}

func isCapital(c int) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c int) bool {
	return c >= 'a' && c <= 'z'
}
