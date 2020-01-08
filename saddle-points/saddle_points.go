package matrix

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MAX = int(^uint(0) >> 1)
	MIN = -MAX - 1
)

type Pair [2]int

type Matrix struct {
	nums [][]int
}

func New(m string) (*Matrix, error) {
	mt := new(Matrix)
	strs := strings.Split(m, "\n")

	mt.nums = make([][]int, len(strs))

	for i, v := range strs {
		fields := strings.Fields(v)
		mt.nums[i] = make([]int, len(fields))

		for j := 0; j < len(fields); j++ {
			t, err := strconv.Atoi(fields[j])
			if err != nil {
				return nil, errors.New("illegal")
			}
			mt.nums[i][j] = t
		}
	}

	return mt, nil
}

func (m *Matrix) Saddle() []Pair {

	tPairs, pairs, nums := []Pair{}, []Pair{}, m.nums

	for i := 0; i < len(nums); i++ { // walk row
		rowMax := MIN
		for j := 0; j < len(nums[0]); j++ { //find the row max
			if nums[i][j] > rowMax {
				rowMax = nums[i][j]
			}
		}

		for j := 0; j < len(nums[0]); j++ {
			if nums[i][j] == rowMax {
				tPairs = append(tPairs, Pair{i, j})
			}
		}

	}

	for i := 0; i < len(nums[0]); i++ { //walk column
		columnMin := MAX
		for j := 0; j < len(nums); j++ { //find the column min
			if nums[j][i] <= columnMin {
				columnMin = nums[j][i]
			}
		}

		for j := 0; j < len(nums); j++ {
			if nums[j][i] == columnMin {
				if isExist(tPairs, j, i) {
					pairs = append(pairs, Pair{j, i})
				}
			}
		}

	}

	return pairs
}

func isExist(pairs []Pair, x int, y int) bool {
	for _, v := range pairs {
		if v[0] == x && v[1] == y {
			return true
		}
	}

	return false
}
