package flatten

func flat(flatted []interface{}, nested interface{}) []interface{} {
	if nested == nil {
		return flatted
	}

	subNested, isSlice := nested.([]interface{});

	if isSlice {
		for _, value := range subNested {
			flatted = flat(flatted, value)
		}
	} else {
		flatted = append(flatted, nested)
	}

	return flatted
}

func Flatten(nested interface{}) []interface{} {
	var flatted = []interface{}{}

	flatted = flat(flatted, nested)

	return flatted
}
