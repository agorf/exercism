package school

import "slices"

type Grade struct {
	level    int
	students []string
}

type School struct {
	grades map[int]Grade
}

func New() *School {
	s := School{}
	s.grades = make(map[int]Grade)
	return &s
}

func (s *School) Add(student string, level int) {
	grade, exists := s.grades[level]
	if !exists {
		grade = Grade{level: level}
	}
	grade.students = append(grade.students, student)
	slices.SortFunc(grade.students, func(a, b string) int {
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	})
	s.grades[level] = grade
}

func (s *School) Grade(level int) []string {
	return s.grades[level].students
}

func (s *School) Enrollment() []Grade {
	var grades []Grade
	for _, g := range s.grades {
		grades = append(grades, g)
	}
	slices.SortFunc(grades, func(a, b Grade) int {
		switch {
		case a.level < b.level:
			return -1
		case a.level > b.level:
			return 1
		default:
			return 0
		}
	})
	return grades
}
