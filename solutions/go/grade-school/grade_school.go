package gradeschool

import (
	"slices"
)

// Grade holds data about a single school grade/year.
type Grade struct {
	level    int
	students []string
}

// School contains data about all the grades.
type School struct {
	grades [20]*Grade
}

// New returns a new School object.
func New() *School {
	return &School{}
}

// Add a student to a specific level/grade.
func (s *School) Add(student string, level int) bool {
	if slices.Contains(s.Enrollment(), student) {
		return false
	}
	grade := s.grades[level]
	if grade == nil {
		// Create a new Grade if it is not yet in the School.
		grade = &Grade{level, []string{}}
		s.grades[level] = grade
	}
	grade.students = append(grade.students, student)
	slices.Sort(grade.students)
	return true
}

// Grade returns the students in a specific Grade.
func (s *School) Grade(level int) []string {
	grade := s.grades[level]
	if grade == nil {
		return nil
	}
	return grade.students
}

// Enrollment returns the students from all Grades.
func (s *School) Enrollment() []string {
	var students []string
	for _, grade := range s.grades {
		if grade != nil {
			students = append(students, grade.students...)
		}
	}
	return students
}
