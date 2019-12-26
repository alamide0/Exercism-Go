package railfence

func Encode(message string, rails int) string {
	arr := make([][]rune, rails)

	for i := 0; i < rails; i++ {
		arr[i] = make([]rune, len(message))
	}

	reverse, row := false, 0

	for i, r := range message {
		arr[row][i] =  r
		if !reverse {
			if row == rails - 1 {
				reverse = true
				row -= 1
			}else {
				reverse = false
				row += 1
			}
		} else {
			if row == 0 {
				reverse = false
				row += 1
			}else {
				reverse = true
				row -= 1
			}
		}
	}

	count := 0
	for i := 0; i < rails; i++ {
		for j := 0; j < len(message); j++ {
			if arr[i][j] != 0 {
				arr[0][count] = arr[i][j]
				count += 1
			}
		}
	}

	return string(arr[0])
} 

func Decode(message string, rails int) string {
	arr := make([][]rune, rails)

	for i := 0; i < rails; i++ {
		arr[i] = make([]rune, len(message))
	}

	reverse, row := false, 0

	// insert tag
	for i := range message {
		arr[row][i] =  'A'
		if !reverse {
			if row == rails - 1 {
				reverse = true
				row -= 1
			}else {
				reverse = false
				row += 1
			}
		} else {
			if row == 0 {
				reverse = false
				row += 1
			}else {
				reverse = true
				row -= 1
			}
		}
	}



	var runes = []rune(message)

	count := 0

	//insert real rune
	for i := 0; i < rails; i++ {
		for j := 0; j < len(message); j++ {
			if arr[i][j] == 'A' {
				arr[i][j] = runes[count]
				count += 1
			}
		}
	}

	//read real message
	row = 0
	reverse = false
	for i := range message {
		runes[i] = arr[row][i]
		if !reverse {
			if row == rails - 1 {
				reverse = true
				row -= 1
			}else {
				reverse = false
				row += 1
			}
		} else {
			if row == 0 {
				reverse = false
				row += 1
			}else {
				reverse = true
				row -= 1
			}
		}
	}

	return string(runes)
} 
