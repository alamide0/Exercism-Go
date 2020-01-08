package school

import (
	"sort"
)

type Grade struct {
	grade int
	names []string
}

type School struct {
	grades []Grade
}

func New() *School {
	return new(School)
}

func (s *School) Enrollment() []Grade {
	return s.grades
}

func (s *School) Add(name string, grade int) {
	needInsert, insertIndex := true, 0

	for i, v := range s.grades {

		if grade > v.grade {
			insertIndex = i + 1
		}

		if v.grade == grade {
			s.grades[i].names = append(s.grades[i].names, name) //not v.names
			sort.Strings(s.grades[i].names)
			needInsert = false
			break
		}
	}

	if needInsert {
		g := Grade{grade, []string{name}}
		s.grades = append(s.grades[:insertIndex], append([]Grade{g}, s.grades[insertIndex:]...)...)
	}
}

func (s *School) Grade(grade int) []string {

	for _, v := range s.grades {
		if v.grade == grade {
			return v.names
		}
	}

	return nil
}
