package flatten

func flat(flatted []interface{}, nested interface{}) []interface{} {
	if nested == nil {
		return flatted
	}

	switch subNested := nested.(type) {
	case []interface{}:
		for _, value := range subNested {
			flatted = flat(flatted, value)
		}
	default:
		flatted = append(flatted, subNested)
	}

	return flatted
}

func Flatten(nested interface{}) []interface{} {
	var flatted = []interface{}{}

	flatted = flat(flatted, nested)

	return flatted
}
