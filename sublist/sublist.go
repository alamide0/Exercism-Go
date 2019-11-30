package sublist

type Relation string

const (
	SUBLIST   = "sublist"
	SUPERLIST = "superlist"
	EQUAL     = "equal"
	UNEQUAL   = "unequal"
)

func Sublist(a, b []int) Relation {
	if len(a) == 0 && len(b) > 0 {
		return SUBLIST
	}

	if len(a) > 0 && len(b) == 0 {
		return SUPERLIST
	}

	if len(a) == 0 && len(b) == 0 {
		return EQUAL
	}

	if len(a) == len(b) {
		for i := range a {
			if a[i] != b[i] {
				return UNEQUAL
			}
		}
		return EQUAL
	} else if len(a) > len(b) {
		for i := 0; i <= len(a)-len(b); i++ {
			if equal(a[i:i+len(b)], b[:]) {
				return SUPERLIST
			}
		}
		return UNEQUAL
	} else {
		for i := 0; i <= len(b)-len(a); i++ {
			if equal(b[i:i+len(a)], a[:]) {
				return SUBLIST
			}
		}
		return UNEQUAL
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
