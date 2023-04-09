package sublist

// Relation type is defined in relations.go file.

func areEqual(l1, l2 []int) Relation {
	for i, v1 := range l1 {
		v2 := l2[i]

		if v1 != v2 {
			return RelationUnequal
		}
	}

	return RelationEqual
}

func isSubList(l1, l2 []int) bool {
	// l2 in l1
	if len(l2) == 0 {
		return true
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] == l2[0] {
			if areEqual(l1[i:i+len(l2)], l2) == RelationEqual {
				return true
			}
		}
	}

	return false
}

func aRelationB(l1, l2 []int, relationAB Relation) Relation {
	if isSubList(l1, l2) {
		return relationAB
	}

	return RelationUnequal
}

func Sublist(l1, l2 []int) Relation {
	var result Relation
	var lenDifference int = len(l1) - len(l2)

	switch {
	case lenDifference == 0:
		result = areEqual(l1, l2)
	case lenDifference < 0:
		result = aRelationB(l2, l1, RelationSublist)
	default:
		result = aRelationB(l1, l2, RelationSuperlist)
	}

	return result
}
