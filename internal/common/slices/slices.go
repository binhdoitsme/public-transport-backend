package slices

func RemoveValue[V comparable](slice []V, value V) []V {
	i := -1
	for idx, item := range slice {
		if item == value {
			i = idx
			break
		}
	}

	if i < 0 {
		return slice
	}

	slice[i] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}
