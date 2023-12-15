package linq

func Aggregate[T any](source []T, initial T, f func(agg, now T) T) T {
	aggregation := initial
	for _, s := range source {
		aggregation = f(aggregation, s)
	}
	return aggregation
}
