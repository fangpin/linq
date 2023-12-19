package linq

func Sum[T Number](source []T) T {
	var sum T
	for _, s := range source {
		sum += s
	}
	return sum
}
