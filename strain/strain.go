package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(fn func(int) bool) Ints {

	var res Ints

	for _, v := range ints {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

func (s Strings) Keep(fn func(string) bool) Strings {
	var res Strings

	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	var res Lists

	for _, v := range lists {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

func (ints Ints) Discard(fn func(int) bool) Ints {
	var res Ints

	for _, v := range ints {
		if !fn(v) {
			res = append(res, v)
		}
	}

	return res
}
