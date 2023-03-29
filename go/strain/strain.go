package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var kept Ints

	if i == nil {
		return nil
	}

	kept = make(Ints, 0, len(i))

	for _, value := range i {
		if filter(value) {
			kept = append(kept, value)
		}
	}

	return kept
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var discarded Ints

	if i == nil {
		return nil
	}

	discarded = make(Ints, 0, len(i))

	for _, value := range i {
		if !filter(value) {
			discarded = append(discarded, value)
		}
	}

	return discarded
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var kept Lists

	if l == nil {
		return nil
	}

	kept = make(Lists, 0, len(l))

	for _, values := range l {
		if filter(values) {
			kept = append(kept, values)
		}
	}

	return kept
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var kept Strings

	if s == nil {
		return nil
	}

	kept = make(Strings, 0, len(s))

	for _, value := range s {
		if filter(value) {
			kept = append(kept, value)
		}
	}

	return kept
}
