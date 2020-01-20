package wordsearch

import (
	"errors"
)

/**
1 1 1
1 2 1
1 1 1

2 is the center
**/
const (
	C1 = iota // (0, 0)
	C2        // (0, 1)
	C3        // (0, 2)
	C4        // (1, 2)
	C5        // (2, 2)
	C6        // (2, 1)
	C7        // (2, 0)
	C8        // (1, 0)
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {

	res := make(map[string][2][2]int)

	bts := make([][]byte, len(puzzle))

	for i := 0; i < len(bts); i++ {
		bts[i] = []byte(puzzle[i])
	}

	for _, v := range words {
		t, ok := findPos(bts, []byte(v))
		if ok {
			res[v] = t
		} else {
			return res, errors.New("fail to locate a word that is not in the puzzle")
		}
	}

	return res, nil

}

func findPos(bts [][]byte, dest []byte) ([2][2]int, bool) {
	for i := 0; i < len(bts); i++ {
		for j := 0; j < len(bts[0]); j++ {
			if bts[i][j] == dest[0] {
				for k := C1; k <= C8; k++ {
					if isExist(bts, dest, 0, i, j, k) {
						endR, endC := caclPos(i, j, k, len(dest)-1)
						return [2][2]int{{j, i}, {endC, endR}}, true
					}
				}

			}
		}
	}

	return [2][2]int{}, false
}

func caclPos(row, col, c, dx int) (destR, destC int) {
	switch c {
	case C1:
		destR, destC = row-dx, col-dx
	case C2:
		destR, destC = row-dx, col
	case C3:
		destR, destC = row-dx, col+dx
	case C4:
		destR, destC = row, col+dx
	case C5:
		destR, destC = row+dx, col+dx
	case C6:
		destR, destC = row+dx, col
	case C7:
		destR, destC = row+dx, col-dx
	case C8:
		destR, destC = row, col-dx
	}
	return
}

func isExist(bts [][]byte, dest []byte, index, row, col, c int) bool {

	if index == len(dest) {
		return true
	}

	if row < 0 || row >= len(bts) || col < 0 || col >= len(bts[0]) {
		return false
	}

	if bts[row][col] != dest[index] {
		return false
	}

	destR, destC := caclPos(row, col, c, 1)

	return isExist(bts, dest, index+1, destR, destC, c)
}
