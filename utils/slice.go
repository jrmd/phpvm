package utils

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

func Some[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if pred(t) {
			return true
		}
	}
	return false
}

func UniqAppend[T comparable](slice []T, item T) []T {
	if Some(slice, func(t T) bool { return t == item }) {
		return slice
	}

	return append(slice, item)
}
