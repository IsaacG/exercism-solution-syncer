package school

import (
	"cmp"
	"slices"
)

type Grade struct {
	level    int
	students []string
}
type School struct {
	grades map[int]*Grade
}

func New() *School {
	return &School{map[int]*Grade{}}
}

func (s *School) Add(student string, level int) {
	grade, ok := s.grades[level]
	if !ok {
		grade = &Grade{level, []string{}}
		s.grades[level] = grade
	}
	grade.students = append(grade.students, student)
	slices.Sort(grade.students)
}

func (s *School) Grade(level int) []string {
	grade, ok := s.grades[level]
	if !ok {
		return nil
	}
	return grade.students
}

func (s *School) Enrollment() []Grade {
	var grades []Grade
	for _, grade := range s.grades {
		grades = append(grades, *grade)
	}
	slices.SortFunc(grades, func(a, b Grade) int { return cmp.Compare(a.level, b.level) })
	return grades
}
