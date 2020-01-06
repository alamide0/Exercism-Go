package transpose

import (
	"bytes"
)

func Transpose(input []string) []string {

	maxLen := 0
	for _, v := range input {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}

	res := make([]string, maxLen)
	var buffer bytes.Buffer

	for i := 0; i < maxLen; i++ {
		trimIndex := -1
		for j, v := range input {
			if i >= len(v) {
				buffer.WriteByte(' ')
			} else {
				buffer.WriteByte(v[i])
				trimIndex = j
			}
		}

		res[i] = buffer.String()[:trimIndex+1]
		buffer.Reset()
	}

	return res
}
