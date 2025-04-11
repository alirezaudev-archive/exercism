package stringset

import (
	"fmt"
	"strings"
)

type Set struct {
	data []string
	size int
}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	var s Set
	for _, v := range l {
		s.Add(v)
	}

	return s
}
func (s Set) String() string {
	for i := range s.data {
		s.data[i] = fmt.Sprintf("\"%s\"", s.data[i])
	}

	return fmt.Sprintf(
		"{%s}",
		strings.Join(s.data, ", "),
	)
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.index(elem)
	return ok
}

func (s *Set) index(elem string) (int, bool) {
	start := 0
	end := s.size - 1
	for start <= end {
		m := (start + end) / 2
		if elem > s.data[m] {
			start = m + 1
		} else if elem < s.data[m] {
			end = m - 1
		} else {
			return m, true
		}
	}

	return start, false
}
func (s *Set) Add(elem string) {
	if start, contains := s.index(elem); !contains {
		s.size++
		s.data = append(s.data, "")
		copy(s.data[start+1:], s.data[start:])
		s.data[start] = elem
	}
}

func Subset(s1, s2 Set) bool {
	if s1.size > s2.size {
		return false
	}

	for i := range s1.data {
		if s1.data[i] != s2.data[i] {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	if s1.size > s2.size {
		s1, s2 = s2, s1
	}
	for i := range s1.data {
		if s2.Has(s1.data[i]) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if s1.size != s2.size {
		return false
	}

	for i := range s1.data {
		if s1.data[i] != s2.data[i] {
			return false
		}
	}

	return true
}

func Intersection(s1, s2 Set) Set {
	var result Set
	for i := range s1.data {
		if s2.Has(s1.data[i]) {
			result.Add(s1.data[i])
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	var result Set
	for i := range s1.data {
		if !s2.Has(s1.data[i]) {
			result.Add(s1.data[i])
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	var result Set
	for i := range s1.data {
		result.Add(s1.data[i])
	}
	for i := range s2.data {
		result.Add(s2.data[i])
	}
	return result
}
