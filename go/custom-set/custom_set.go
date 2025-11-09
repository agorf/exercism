package stringset

import "strings"

type seen struct{}

type Set map[string]seen

var seenValue = seen{}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := Set{}
	for _, r := range l {
		set[string(r)] = seenValue
	}
	return set
}

func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}
	var elems []string
	for elem := range s {
		elems = append(elems, `"`+elem+`"`)
	}
	return "{" + strings.Join(elems, ", ") + "}"
}

func (s Set) Length() int {
	return len(s)
}

func (s Set) IsEmpty() bool {
	return s.Length() == 0
}

func (s Set) Has(elem string) bool {
	_, exists := s[elem]
	return exists
}

func (s Set) Add(elem string) {
	s[elem] = seenValue
}

func Subset(s1, s2 Set) bool {
	if s1.Length() > s2.Length() {
		return false
	}
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if s1.Length() != s2.Length() {
		return false
	}
	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	a := s1
	b := s2
	if a.Length() > b.Length() {
		a, b = b, a
	}
	set := Set{}
	for elem := range a { // Loop over the smallest set
		if b.Has(elem) {
			set[elem] = seenValue
		}
	}
	return set
}

func Difference(s1, s2 Set) Set {
	set := Set{}
	for elem := range s1 {
		if !s2.Has(elem) {
			set[elem] = seenValue
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	set := Set{}
	for elem := range s1 {
		set[elem] = seenValue
	}
	for elem := range s2 {
		set[elem] = seenValue
	}
	return set
}
