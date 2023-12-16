package linq

func GroupJoin[T, R, S any, K comparable](
	outter []T,
	inner []R,
	outterKeySelector func(T) K,
	innerKeySelector func(R) K,
	resultSelector func(T, []R) S) []S {

	innerMap := map[K][]R{}
	for _, s := range inner {
		key := innerKeySelector(s)
		list := innerMap[key]
		list = append(list, s)
		innerMap[key] = list
	}
	result := []S{}
	for _, s := range outter {
		key := outterKeySelector(s)
		inners := innerMap[key]
		result = append(result, resultSelector(s, inners))
	}
	return result
}
