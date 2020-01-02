package diamond
import (
	"errors"
	"strings"
)

func Gen(b byte) (string, error) {

	if b > 'Z' || b < 'A' {
		return "", errors.New("illegal char")
	}

	if b == 'A' {
		return "A\n", nil
	}

	n := int(b - 'A') + 1
	lines := int(2 * n - 1)

	insertP := n

	var builder strings.Builder

	for i := 1; i <= n; i++ {
		for j := 1; j <= lines; j++ {
			if j == insertP || j == lines - insertP + 1{
				builder.WriteRune(rune('A' + i - 1))
			}else {
				builder.WriteRune(rune(' '))
			}
		}
		builder.WriteRune('\n')
		insertP -= 1
	}

	strs := strings.SplitAfter(builder.String(), "\n")

	for i := len(strs) - 3; i >=0; i-- {
		builder.WriteString(strs[i])
	}

	return builder.String(), nil
}
