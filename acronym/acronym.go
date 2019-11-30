
package acronym

func Abbreviate(s string) string {
	
	isNew := true
	var runes []rune

	for _, r := range s {
		if isLetter, c := isLetter(r); (isNew && isLetter) {
			isNew = false
			runes = append(runes, c)
		} else if r == '-' || r == ' '{
			isNew = true
		}
	}

	return string(runes)
}

func isLetter(r rune) (bool, rune) {
	if (r >= 'A' && r <= 'Z') {
		return true, r
	}

	if (r >= 'a' && r <= 'z') {
		return true, r + ('A' - 'a') 
	}

	return false, r
}


