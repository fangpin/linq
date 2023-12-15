package linq

func Except[T comparable](l []T, r []T) []T {
	setR := ToHashSet(r)
	target := []T{}
	for _, x := range l {
		if _, ok := setR[x]; !ok {
			target = append(target, x)
		}
	}
	return target
}
