package matrix
import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

const INT64_MAX = int(^uint64(0) >> 1)

type Matrix struct {
	data []int
	width int
}
/*
0 0 0
0 0 0
width = 3
height = 2
*/

func New(input string)(*Matrix, error){

	if strings.HasSuffix(input, "\n"){
		return nil, fmt.Errorf("last row empty")
	}

	buf := bufio.NewReader(strings.NewReader(input))

	m := &Matrix{}

	var line string
	for {
		bs, isPrefix, err := buf.ReadLine()
		if err != nil {
			break
		}

		line += string(bs)

		if isPrefix {
			continue
		}
			
		line = strings.TrimSpace(line)
		if line == "" {
			return nil, fmt.Errorf("row empty")
		}

		strs := strings.Fields(line)

		if m.width != 0 && m.width != len(strs) {
			return nil, fmt.Errorf("uneven rows")
		}
		
		m.width = len(strs)

		for _, v := range strs {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			m.data = append(m.data, n)
		}

		line = ""
		
	}

	return m, nil
}

func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r > len(m.data)/m.width - 1 || c > m.width - 1 || val > INT64_MAX{
		return false
	}

	m.data[r*m.width + c] = val
	return true
}

func (m *Matrix) Rows() [][]int {
	w := m.width
	h := len(m.data) / w
	arr := make([][]int, h)
	for i := 0; i < h; i++ {
		arr[i] = make([]int, w)
		for j := 0; j < w; j++ {
			arr[i][j] = m.data[i*w + j]
		}
	}
	return arr
}

func (m *Matrix) Cols() [][]int {
	h := m.width
	w := len(m.data) / h
	arr := make([][]int, h)
	for i := 0; i < h; i++ {
		arr[i] = make([]int, w)
		for j := 0; j < w; j++ {
			arr[i][j] = m.data[i + j*h]
		}
	}
	return arr
}