package stringset

import (
	"strings"
)

type Set map[string]bool

func (set Set) Len() int {
	return len(set)
}

func (set Set) String() string {
	var builder strings.Builder
	builder.WriteRune('{')
	count := 0
	for k, _ := range set {
		builder.WriteString("\"")
		builder.WriteString(k)
		builder.WriteString("\"")
		if count != len(set)-1 {
			builder.WriteString(", ")
		}
		count++
	}

	builder.WriteRune('}')

	return builder.String()
}

func New() Set {
	return make(Set)
}

func NewFromSlice(strs []string) Set {
	set := New()
	for _, v := range strs {
		set.Add(v)
	}
	return set
}

func (set Set) Has(str string) bool {
	_, ok := set[str]
	return ok
}

func (set Set) Add(str string) {
	set[str] = true
}

func (set Set) IsEmpty() bool {
	return len(set) == 0
}

func Subset(s1, s2 Set) bool {
	for k, _ := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for k, _ := range s1 {
		if s2.Has(k) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {

	if len(s1) != len(s2) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {

	set := New()

	for k, _ := range s1 {
		if s2.Has(k) {
			set.Add(k)
		}
	}

	return set
}

func Difference(s1, s2 Set) Set {
	set := New()

	for k, _ := range s1 {
		if !s2.Has(k) {
			set.Add(k)
		}
	}

	return set
}

func Union(s1, s2 Set) Set {
	for k, _ := range s1 {
		s2.Add(k)
	}

	return s2
}
