package utils

func Contains[T comparable](items []T, item T) bool {
	for _, t := range items {
		if t == item {
			return true
		}
	}

	return false
}

func Map[T, D any](ts []T, mapper func(t T) D) []D {
	ds := make([]D, len(ts))

	for i, t := range ts {
		ds[i] = mapper(t)
	}

	return ds
}
