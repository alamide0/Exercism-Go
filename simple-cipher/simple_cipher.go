package cipher

type Key struct {
	shifts []int
}

func NewCaesar() Cipher {
	return &Key{[]int{3}}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return Cipher(nil)
	}
	return &Key{[]int{distance}}
}

func NewVigenere(key string) Cipher {

	k := &Key{make([]int, len(key))}

	isAlla := true

	for i, r := range key {
		if r < 'a' || r > 'z' {
			return Cipher(nil)
		}

		if r != 'a' {
			isAlla = false
		}
		k.shifts[i] = int(r - 'a')
	}

	if isAlla {
		return Cipher(nil)
	}

	return k

}

func (k *Key) Encode(text string) string {
	return deal(k, text, true)
}

func (k *Key) Decode(text string) string {
	return deal(k, text, false)
}

func deal(k *Key, text string, isEncode bool) string {
	runes := make([]rune, 0, len(text))

	index := 0

	for _, r := range text {

		isExist := false
		if r >= 'a' && r <= 'z' {
			isExist = true
		} else if r >= 'A' && r <= 'Z' {
			isExist = true
			r += 32
		}

		if !isExist {
			continue
		}

		t := 'a'
		if isEncode {
			t = rune(int(r) + k.shifts[index])
		} else {
			t = rune(int(r) - k.shifts[index])
		}

		index = index + 1
		if index == len(k.shifts) {
			index = 0
		}

		if t < 'a' {
			t = 'z' - ('a' - t - 1)
		} else if t > 'z' {
			t = 'a' + (t - 'z' - 1)
		}

		runes = append(runes, t)

	}

	return string(runes)
}
