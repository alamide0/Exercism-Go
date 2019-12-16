package etl

func Transform(table map[int][]string) map[string]int {

	res := make(map[string]int)

	for k, v := range table {
		for _, r := range v {
			res[string('a' + r[0] - 'A')] = k
		}
	}

	return res
}
