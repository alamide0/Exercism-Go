package accumulate

func Accumulate(given []string, converter func(s string) string) []string {
	for i, v := range given {
		given[i] = converter(v)
	}
	return given
}