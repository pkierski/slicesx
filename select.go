package slicesx

// Select selects elements from a slice that satisfy a predicate.
func Select[S ~[]E, E any](slice S, predicate func(E) bool) S {
	result := make(S, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
