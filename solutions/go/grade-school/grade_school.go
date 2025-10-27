package school

import (
	"cmp"
	"slices"
)

// Grade holds data about a single school grade/year.
type Grade struct {
	level    int
	students []string
}

// School contains data about all the grades.
type School struct {
	grades map[int]*Grade
}

// New returns a new School object.
func New() *School {
	return &School{map[int]*Grade{}}
}

// Add a student to a specific level/grade.
func (s *School) Add(student string, level int) {
	grade, ok := s.grades[level]
	if !ok {
		// Create a new Grade if it is not yet in the School.
		grade = &Grade{level, []string{}}
		s.grades[level] = grade
	}
	grade.students = append(grade.students, student)
	slices.Sort(grade.students)
}

// Grade returns the students in a specific Grade.
func (s *School) Grade(level int) []string {
	grade, ok := s.grades[level]
	if !ok {
		return nil
	}
	return grade.students
}

// Enrollment returns the students from all Grades.
func (s *School) Enrollment() []Grade {
	var grades []Grade
	for _, grade := range s.grades {
		grades = append(grades, *grade)
	}
	slices.SortFunc(grades, func(a, b Grade) int { return cmp.Compare(a.level, b.level) })
	return grades
}
