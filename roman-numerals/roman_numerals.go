package romannumerals

import (
	"errors"
	"strings"
)

var pairs = map[int]string{
	900: "CM",
	500: "D",
	400: "CD",
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I",
}

func ToRomanNumeral(arabic int) (string, error) {

	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("illegal number")
	}

	n := 0

	if arabic >= 1000 {
		n = 4
	} else if arabic >= 100 {
		n = 3
	} else if arabic >= 10 {
		n = 2
	} else {
		n = 1
	}

	arr := []int{
		0,
		1,
		10,
		100,
		1000,
	}

	var res strings.Builder

	for i := n; i > 0; i-- {
		a := arabic / arr[i]
		arabic = arabic % arr[i]
		switch i {
		case 4:
			for j := 0; j < a; j++ {
				res.WriteString("M")
			}
		case 3, 2, 1:
			if a == 9 || a == 4 {
				res.WriteString(pairs[a*arr[i]])
			} else if a >= 5 {
				res.WriteString(pairs[5*arr[i]])
				for j := 0; j < a-5; j++ {
					res.WriteString(pairs[arr[i]])
				}
			} else {
				for j := 0; j < a; j++ {
					res.WriteString(pairs[arr[i]])
				}
			}
		}
	}

	return res.String(), nil
}
