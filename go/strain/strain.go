package strain

func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range collection {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	for _, v := range collection {
		if !predicate(v) {
			result = append(result, v)
		}
	}

	return result
}
