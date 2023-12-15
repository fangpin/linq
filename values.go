package linq

func Values[T comparable, R any](source map[T]R) []R {
	target := []R{}
	for _, v := range source {
		target = append(target, v)	
	}	
	return target
}