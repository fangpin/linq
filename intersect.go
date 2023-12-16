package linq

func Intersect[T comparable](l, r []T) []T {
	set := ToHashSet(l)
	target := []T{}
	for _, s := range r {
		if _, ok := set[s]; ok {
			target = append(target, s)
		}
	}
	return target
}
