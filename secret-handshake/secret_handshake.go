package secret

func Handshake(code uint)(result []string) {
	var result []string
	if code & 1 != 0 {
		result = append(result, "wink")
	}

	if code & 2 != 0 {
		result = append(result, "double blink")
	}

	if code & 4 != 0 {
		result = append(result, "close your eyes")
	}

	if code & 8 != 0 {
		result = append(result, "jump")
	}

	if code & 16 != 0 {
		for i, j := 0, len(result) - 1; i < j; i, j = i+1, j -1 {//rervese
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}