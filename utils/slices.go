package utils

func Contains[T comparable](items []T, item T) bool {
	for _, t := range items {
		if t == item {
			return true
		}
	}

	return false
}
