package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	var output Set = make(Set, len(l))

	for _, s := range l {
		output.Add(s)
	}

	return output
}

func (s Set) String() string {
	var output strings.Builder
	var counter int

	output.WriteRune('{')

	for str := range s {
		output.WriteRune('"')
		output.WriteString(str)
		output.WriteRune('"')

		counter++

		if counter < len(s) {
			output.WriteString(", ")
		}
	}

	output.WriteRune('}')

	return output.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}

	for elem := range s1 {
		if exists := s2[elem]; !exists {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	var output Set = New()

	for elem := range s1 {
		if s2[elem] {
			output.Add(elem)
		}
	}

	return output
}

func Difference(s1, s2 Set) Set {
	var output Set = New()
	var intersection Set = Intersection(s1, s2)

	for elem := range s1 {
		if !intersection.Has(elem) {
			output.Add(elem)
		}
	}

	return output
}

func Union(s1, s2 Set) Set {
	var output Set = New()

	for elem := range s1 {
		output.Add(elem)
	}

	for elem := range s2 {
		output.Add(elem)
	}

	return output
}
