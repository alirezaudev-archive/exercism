package school

import "sort"

type School struct {
	students map[string]int
	grades   map[int]map[string]bool
}

type Grade struct {
	Level    int
	Students []string
}

func New() *School {
	return &School{
		students: make(map[string]int),
		grades:   make(map[int]map[string]bool),
	}
}

func (s *School) Add(student string, grade int) {
	if _, exists := s.students[student]; exists {
		return
	}

	s.students[student] = grade

	if s.grades[grade] == nil {
		s.grades[grade] = make(map[string]bool)
	}
	s.grades[grade][student] = true
}

func (s *School) Grade(level int) []string {
	studentMap := s.grades[level]
	if studentMap == nil {
		return []string{}
	}

	students := make([]string, 0, len(studentMap))
	for student := range studentMap {
		students = append(students, student)
	}

	sort.Strings(students)
	return students
}

func (s *School) Enrollment() []Grade {
	levels := make([]int, 0, len(s.grades))
	for level := range s.grades {
		levels = append(levels, level)
	}
	sort.Ints(levels)

	result := make([]Grade, 0, len(levels))
	for _, level := range levels {
		students := s.Grade(level)
		result = append(result, Grade{
			Level:    level,
			Students: students,
		})
	}

	return result
}
