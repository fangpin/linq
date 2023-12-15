package linq

func Average[T Number](source []T) float64 {
	var ans float64
	for _, s := range source {
		ans += float64(s)
	}
	return ans / float64(len(source))
}
