package minesweeper

import (
	"errors"
)

func (b Board) Count() error {
	if isIllegal(b) {
		return errors.New("illegal")
	}

	for i := 1; i < len(b)-1; i++ {
		for j := 1; j < len(b[0])-1; j++ {
			if b[i][j] == '*' {
				continue
			}
			count := 0
			for k := i - 1; k <= i+1; k++ {
				for h := j - 1; h <= j+1; h++ {
					if b[k][h] == '*' {
						count++
					}
				}
			}
			if count != 0 {
				b[i][j] = '0' + byte(count)
			}
		}
	}

	return nil

}

func isIllegal(b Board) bool {
	if len(b) < 2 || len(b[0]) < 2 {
		return true
	}

	w, h := len(b[0]), len(b)

	for i := 1; i < len(b); i++ {
		if w != len(b[i]) {
			return true
		}
	}

	if b[0][0] != '+' || b[0][w-1] != '+' || b[h-1][0] != '+' || b[h-1][w-1] != '+' {
		return true
	}

	for i := 1; i < w-1; i++ {
		if b[0][i] != '-' || b[h-1][i] != '-' {
			return true
		}
	}

	for i := 1; i < h-1; i++ {
		if b[i][0] != '|' || b[i][w-1] != '|' {
			return true
		}
	}

	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			if b[i][j] != '*' && b[i][j] != ' ' {
				return true
			}
		}
	}

	return false
}