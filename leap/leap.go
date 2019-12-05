// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	
	if year % 4 != 0 {
		return false
	}

	if year % 400 == 0 {
		return true
	}

	if year % 100 == 0 {
		return false
	}

	return true
}
//above Benchmark400-8           3000000               425 ns/op               0 B/op          0 allocs/op
//faster than return (year % 4 == 0 && year % 100 != 0) || year % 400 == 0
/**
and faster than 
Benchmark400-8           2000000               877 ns/op               0 B/op          0 allocs/op

if year % 4 != 0 {
	return false
}

return year % 100 != 0 || year % 400 == 0
???
somebody can tell me why?
*/
