package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plantsMap = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

type Garden struct {
	g map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	strs := strings.Split(diagram, "\n")
	strs = strs[1:]
	if len(strs) != 2 || len(strs[0]) != len(strs[1]) || len(strs[0])%2 != 0 {
		return nil, errors.New("illegal ")
	}

	sorted := make([]string, len(children))

	copy(sorted, children)

	res := &Garden{g: make(map[string][]string)}

	sort.Strings(sorted)

	for i, v := range sorted {
		if _, ok := res.g[v]; ok {
			return nil, errors.New("duplicate name")
		}

		res.g[v] = make([]string, 4)
		t := res.g[v]

		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				if m, ok := plantsMap[strs[j][i*2+k]]; ok {
					t[j*2+k] = m
				} else {
					return nil, errors.New("invaid cup codes")
				}
			}
		}
	}

	return res, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	res, ok := g.g[child]
	if !ok {
		return nil, false
	}

	return res, true
}
