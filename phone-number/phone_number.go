package phonenumber

import (
	"errors"
)

func Number(input string) (string, error) {

	phone := make([]rune, 11)
	count := 0
	for _, r := range input {
		if r >= '0' && r <= '9' {
			if count >= 11 {
				return "", errors.New("invaild phone")
			}
			phone[count] = r
			count++
		}
	}

	if count == 11 && phone[0] == '1' {
		phone = phone[1:]
	} else if count == 10 {
		phone = phone[:10]
	} else {
		return "", errors.New("invaild phone")
	}

	if phone[0] < '2' || phone[3] < '2' {
		return "", errors.New("invaild phone")
	}

	return string(phone), nil
}

func AreaCode(input string) (string, error) {
	phone, err := Number(input)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

func Format(input string) (string, error) {
	phone, err := Number(input)
	if err != nil {
		return "", err
	}
	return "(" + phone[:3] + ") " + phone[3:6] + "-" + phone[6:], nil
}
