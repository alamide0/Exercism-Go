package reverse

func Reverse(input string) string {
	runes := []rune(input)
	length, tag := len(runes), len(runes) / 2

	for i := 0; i < tag; i++ {
		runes[i], runes[length - 1 - i] = runes[length - 1 - i], runes[i]
	}
	return string(runes)
}
