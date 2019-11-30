package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	max = 26 * 26 * 10 * 10 * 10
)

var robots = make(map[Robot]bool)

type Robot string

func (r *Robot) Name() (string, error){

	if *r != "" {
		return string(*r), nil
	}

	if len(robots) > max {
		return "", errors.New("no more names")
	}

	rand.Seed(time.Now().UnixNano())
	var name string

	for {
		name = ""
		name += fmt.Sprintf("%c%c%03d", ('A' + rand.Intn(26)), ('A' + rand.Intn(26)), rand.Intn(1000))
		if !robots[Robot(name)]{
			*r = Robot(name)
			robots[*r] = true
			break
		}
	}

	return name, nil
}

func (r *Robot) Reset() {
	*r = ""
}