package brackets

/**
use stack
({}{}{}{}{}{}{})
({}{}{}{}{}{}{})
({}{}{}{}{}{})
({}{}{}{}{}{})
({}{}{}{}{})
({}{}{}{}{})
({}{}{}{})
({}{}{}{})
({}{}{})
({}{}{})
({}{})
({}{})
({})
({})
()

last will be empty
**/
func Bracket(str string) bool {
	var runes []rune
	for _, r := range str {
		if r == '{' || r == '}' || r == '(' || r == ')' || r == '[' || r == ']' {
			runes = append(runes, r)
		}
	}

	if len(runes) % 2 != 0 {
		return false
	}

	for i := 0; i < len(runes);{
		switch runes[i] {
		case ')':
			if i - 1 < 0 || runes[i-1] != '(' {
				return false
			}
			runes = append(runes[:i - 1], runes[i+1:]...)
			i -= 1
		case '}':
			if i - 1 < 0 || runes[i-1] != '{' {
				return false
			}
			runes = append(runes[:i - 1], runes[i+1:]...)
			i -= 1
		case ']':
			if i - 1 < 0 || runes[i-1] != '[' {
				return false
			}
			runes = append(runes[:i - 1], runes[i+1:]...)
			i  -= 1
		default:
			i++
		}
	}

	return len(runes) == 0
}



