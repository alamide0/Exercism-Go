package stringset

import (
	"strings"
)

type Set struct {
	strs []string
}

func (set Set) Len() int {
	return len(set.strs)
}

func (set Set) String() string {
	var builder strings.Builder
	builder.WriteRune('{')
	for i := 0; i < set.Len(); i++ {
		builder.WriteString("\"")
		builder.WriteString(set.strs[i])
		builder.WriteString("\"")
		if i != set.Len()-1 {
			builder.WriteString(", ")
		}
	}

	builder.WriteRune('}')

	return builder.String()
}

func New() Set {
	return *new(Set)
}

func NewFromSlice(strs []string) Set {
	set := new(Set)
	for _, v := range strs {
		set.Add(v)
	}
	return *set
}

func (set *Set) Has(str string) bool {
	for _, v := range set.strs {
		if v == str {
			return true
		}
	}
	return false
}

func (set *Set) Add(str string) {
	if !set.Has(str) {
		set.strs = append(set.strs, str)
	}
}

func (set *Set) IsEmpty() bool {
	return set.Len() == 0
}

func Subset(s1, s2 Set) bool {
	for _, v := range s1.strs {
		if !s2.Has(v) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, v := range s1.strs {
		if s2.Has(v) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {

	if s1.Len() != s2.Len() {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {

	set := new(Set)

	for _, v := range s1.strs {
		if s2.Has(v) {
			set.Add(v)
		}
	}

	return *set
}

func Difference(s1, s2 Set) Set {
	set := new(Set)

	for _, v := range s1.strs {
		if !s2.Has(v) {
			set.Add(v)
		}
	}

	return *set
}

func Union(s1, s2 Set) Set {
	for _, v := range s1.strs {
		(&s2).Add(v)
	}

	return s2
}
