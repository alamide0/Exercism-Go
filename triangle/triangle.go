// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle
import (
	"math"
)

type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = iota
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a + b + c) || math.IsInf(a + b + c, 0) ||  a <=0 || b <= 0 || c <= 0 || a + b < c || a + c < b || b + c < a {//is NaT?
		return NaT
	}

	if a == b && a == c {//is equilateral?
		return Equ
	}

	if a == b || a == c || b == c {//is isosceles?
		return Iso
	}

	return Sca
}
