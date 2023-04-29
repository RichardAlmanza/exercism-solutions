package flatten

func flat(flatted *[]interface{}, nested interface{}) {
	if nested == nil {
		return
	}

	switch subNested := nested.(type) {
	case []interface{}:
		for _, value := range subNested {
			flat(flatted, value)
		}
	default:
		*flatted = append(*flatted, subNested)
	}
}

func Flatten(nested interface{}) []interface{} {
	var flatted = []interface{}{}

	flat(&flatted, nested)

	return flatted
}
