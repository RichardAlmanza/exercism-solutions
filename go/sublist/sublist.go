package sublist

// Relation type is defined in relations.go file.

type listComparator func(a,b []int)bool

func areEqual(l1, l2 []int) bool {
	for i, v1 := range l1 {
		v2 := l2[i]

		if v1 != v2 {
			return false
		}
	}

	return true
}

func isSubList(l1, l2 []int) bool {
	// l2 in l1
	if len(l2) == 0 {
		return true
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] == l2[0] {
			if areEqual(l1[i:i+len(l2)], l2) {
				return true
			}
		}
	}

	return false
}

func aRelationB(l1, l2 []int, relationAB Relation, comparator listComparator) Relation {
	if comparator(l1, l2) {
		return relationAB
	}

	return RelationUnequal
}

func Sublist(l1, l2 []int) Relation {
	var result Relation

	switch {
	case len(l1) == len(l2):
		result = aRelationB(l1, l2, RelationEqual, areEqual)
	case len(l1) < len(l2):
		result = aRelationB(l2, l1, RelationSublist, isSubList)
	default:
		result = aRelationB(l1, l2, RelationSuperlist, isSubList)
	}

	return result
}
