package tournament
import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name string
	win int
	loss int
	draw int
	point int
	matches int
}

func Tally (r io.Reader, w io.Writer) error {

	m := make(map[string]*Team)

	buf := bufio.NewReader(r)

	var line string
	var strs []string
	for {
		bts, _, err := buf.ReadLine()
		if err != nil {
			break
		}

		line = string(bts)
		if line == "" || strings.HasPrefix(line, "#"){
			continue
		}

		strs = strings.Split(line, ";")

		if len(strs) != 3 {
			return fmt.Errorf("error format")
		}

		if _, ok := m[strs[0]]; !ok {
			m[strs[0]] = &Team{name: strs[0]}
		}

		if _, ok := m[strs[1]]; !ok {
			m[strs[1]] = &Team{name: strs[1]}
		}

		switch strs[2] {
		case "win":
			m[strs[0]].win += 1
			m[strs[1]].loss += 1
		case "loss":
			m[strs[0]].loss += 1
			m[strs[1]].win += 1
		case "draw":
			m[strs[0]].draw += 1
			m[strs[1]].draw += 1
		default:
			return fmt.Errorf("no match result")
		}
		m[strs[0]].matches += 1
		m[strs[1]].matches += 1

	}
	
	var teams []*Team
	for _, v := range m {
		v.point = v.win * 3 + v.draw 
		teams = append(teams, v)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].point == teams[j].point {
			return teams[i].name < teams[j].name
		}
		return teams[i].point > teams[j].point
	})

	writeHeader(w)
	for _, v := range teams {
		w.Write([]byte(fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n", v.name, v.matches, v.win, v.draw, v.loss, v.point)))
	}

	return nil
}

func writeHeader(w io.Writer){
	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
}