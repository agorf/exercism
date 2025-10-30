package strain

func Keep[T any](s []T, predicate func(T) bool) []T {
	kept := make([]T, 0)
	for _, x := range s {
		if predicate(x) {
			kept = append(kept, x)
		}
	}
	return kept
}

func Discard[T any](s []T, predicate func(T) bool) []T {
	return Keep(s, func(x T) bool {
		return !predicate(x)
	})
}
