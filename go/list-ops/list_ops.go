package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	var sLen int = s.Length()

	for i := 0; i < sLen; i++ {
		initial = fn(initial, s[i])
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	var sLen int = s.Length()

	for i := sLen - 1; i > -1; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var counter int
	var filtered IntList
	var newList IntList = make(IntList, s.Length())

	for _, value := range s {
		if fn(value) {
			newList[counter] = value
			counter++
		}
	}

	filtered = make(IntList, counter)

	for i := 0; i < counter; i++ {
		filtered[i] = newList[i]
	}

	return filtered
}

func (s IntList) Length() int {
	var counter int

	for range s {
		counter++
	}

	return counter
}

func (s IntList) Map(fn func(int) int) IntList {
	var newList IntList = make(IntList, s.Length())

	for i, value := range s {
		newList[i] = fn(value)
	}

	return newList
}

func (s IntList) Reverse() IntList {
	var sLen int = s.Length()
	var newList IntList = make(IntList, sLen)

	for i := 0; i < sLen; i++ {
		newList[i] = s[sLen - 1 - i]
	}

	return newList
}

func (s IntList) Append(lst IntList) IntList {
	var sLen int = s.Length()
	var lstLen int = lst.Length()
	var newListLen int = sLen + lstLen
	var newList IntList = make(IntList, newListLen)

	for i, value := range s {
		newList[i] = value
	}

	for i, value := range lst {
		newList[sLen + i] = value
	}

	return newList
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}

	return s
}
