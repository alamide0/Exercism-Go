package anagram

import (
	"sort"
	"strings"
)

type Runes []rune

func (r Runes) Len() int {
	return len(r)
}

func (r Runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Runes) Less(i, j int) bool {
	return r[i] < r[j]
}

func Detect(sub string, candidates []string) (res []string) {
	sub = strings.ToLower(sub)
	runes := Runes(sub)
	sort.Sort(runes)

	for _, v := range candidates {
		s := strings.ToLower(v)
		if s == sub {
			continue
		}
		rs := Runes(s)
		sort.Sort(rs)
		if isEqual(runes, rs) {
			res = append(res, v)
		}
	}

	return
}

func isEqual(r1, r2 Runes) bool {

	if r1.Len() != r2.Len() {
		return false
	}

	for i := 0; i < r1.Len(); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}
