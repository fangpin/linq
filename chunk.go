package linq

func Chunk[T any](source []T, size int) [][]T {
	target := [][]T{}
	cnt := 0
	chunk := []T{}
	for _, s := range source {
		chunk = append(chunk, s)
		cnt++
		if cnt == size {
			target = append(target, chunk)
			chunk = []T{}
			cnt = 0
		}
	}
	if len(chunk) > 0 {
		target = append(target, chunk)
	}
	return target
}
