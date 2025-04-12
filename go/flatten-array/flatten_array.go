package flatten

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}

	if nested == nil {
		return result
	}

	if slice, ok := nested.([]interface{}); ok {
		for _, item := range slice {
			flattened := Flatten(item)
			result = append(result, flattened...)
		}
	} else {
		result = append(result, nested)
	}

	return result
}
