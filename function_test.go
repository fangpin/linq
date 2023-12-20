package linq

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"
)

var emptySource []int

func TestLast(t *testing.T) {
	source := []int{1, 3, 5, 7, 8}
	target := Last(source, func(x int) bool { return x%2 == 1 })
	assert.Equal(t, target, 7)
}

func TestLastPanic(t *testing.T) {
	source := []int{1, 3, 5, 7}
	assert.Panic(t, func() { Last(source, func(x int) bool { return x%2 == 0 }) }, ElementNotFound)
}

func TestLastOrDefault(t *testing.T) {
	source := []int{1, 3, 5, 7, 8}
	target := LastOrDefault(source, func(x int) bool { return x%2 == 1 })
	assert.Equal(t, target, 7)

	source = []int{1, 3, 5, 7}
	target = LastOrDefault(source, func(x int) bool { return x%2 == 0 })
	assert.Equal(t, target, 0)
}

func TestMax(t *testing.T) {
	source := []int{1, 3, 5, 0, -1}
	target := Max(source)
	assert.Equal(t, target, 5)
	var emptySource []int
	assert.Panic(t, func() { Max(emptySource) }, EmptySource)
}

func TestMaxBy(t *testing.T) {
	source := []int{1, 3, 5, 0}
	target := MaxBy(source, func(x int) int { return x % 3 })
	assert.Equal(t, target, 5)
	assert.Panic(
		t,
		func() { MaxBy(emptySource, func(x int) int { return x }) },
		EmptySource)
}

func TestMin(t *testing.T) {
	source := []int{1, 3, 5, 0}
	assert.Equal(t, Min(source), 0)
	var emptySource []int
	assert.Panic(t, func() { Min(emptySource) }, EmptySource)
}

func TestMinBy(t *testing.T) {
	source := []int{1, 3, 5, 7}
	assert.Equal(t, MinBy(source, func(x int) int { return x % 3 }), 3)
	assert.Panic(t, func() { MinBy(emptySource, func(x int) int { return x % 3 }) }, EmptySource)
}

func TestOfType(t *testing.T) {
	source := []any{1, 3, 2.8}
	assert.Equal(t, len(OfType(source, reflect.TypeOf(1))), 2)
}

func TestOrderDescending(t *testing.T) {
	source := []int{1, 3, 2, 9, 6}
	assert.Equal(t, OrderDescending(source), []int{9, 6, 3, 2, 1})
}

func TestOrderByDescending(t *testing.T) {
	source := []int{4, 6, 5}
	assert.Equal(t, OrderByDescending(source, func(l, r int) bool { return (l % 3) < (r % 3) }), []int{5, 4, 6})
}

func TestOrderBy(t *testing.T) {
	source := []int{4, 2, 7, 1}
	assert.Equal(t, OrderBy(source, func(l, r int) bool { return l%4 < r%4 }), []int{4, 1, 2, 7})
}

func TestOrder(t *testing.T) {
	source := []int{1, 3, 5, 2, 4, 6}
	assert.Equal(t, Order(source), []int{1, 2, 3, 4, 5, 6})
}

func TestPrepend(t *testing.T) {
	source := []int{1, 2, 3}
	assert.Equal(t, Prepend(source, 4, 5, 6), []int{6, 5, 4, 1, 2, 3})
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, Repeat(2, 4), []int{2, 2, 2, 2})
}

func TestReverse(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, Reverse(source), []int{6, 5, 4, 3, 2, 1})
}

func TestSelectMany(t *testing.T) {
	source := []int{1, 3}
	assert.Equal(t, SelectMany(source, func(x int) []string {
		return []string{strconv.Itoa(x), strconv.Itoa(x * 2)}
	}), []string{"1", "2", "3", "6"})
}

func TestSingle(t *testing.T) {
	source := []int{1, 3, 4, 5, 7}
	assert.Equal(t, Single(source, func(x int) bool {
		return x%2 == 0
	}), 4)
	assert.Panic(
		t,
		func() {
			Single([]int{1, 3, 5, 7, 9}, func(x int) bool {
				return x%2 == 0
			})
		},
		ElementNotFound)
	assert.Panic(
		t,
		func() {
			Single([]int{1, 3, 6, 8, 9}, func(x int) bool {
				return x%2 == 0
			})
		},
		MultipleElementFound)
}

func TestSingleOrDefault(t *testing.T) {
	assert.Equal(
		t,
		SingleOrDefault([]int{1, 3, 4, 5, 7}, func(x int) bool {
			return x%2 == 0
		}),
		4)
	assert.Equal(
		t,
		SingleOrDefault([]int{1, 3, 9, 5, 7}, func(x int) bool {
			return x%2 == 0
		}),
		0)
	assert.Panic(
		t,
		func() {
			SingleOrDefault([]int{1, 3, 5, 7, 9}, func(x int) bool {
				return x%2 == 1
			})
		},
		MultipleElementFound)
}

func TestSkip(t *testing.T) {
	assert.Equal(
		t,
		Skip([]int{1, 2, 3, 4}, 2),
		[]int{3, 4})
	assert.Panic(
		t,
		func() {
			Skip(emptySource, 2)
		},
		InvalidParam)
}

func TestSkipLast(t *testing.T) {
	assert.Equal(
		t,
		SkipLast([]int{1, 2, 3, 4}, 2),
		[]int{1, 2})
	assert.Panic(
		t,
		func() { SkipLast([]int{1, 2, 3, 4}, 7) },
		InvalidParam)
	assert.Panic(
		t,
		func() { SkipLast([]int{1, 2, 3, 4}, -1) },
		InvalidParam)
}

func TestSkipWhile(t *testing.T) {
	assert.Equal(
		t,
		SkipWhile([]int{1, 2, 3, 4}, func(x int) bool {
			return x <= 3
		}),
		[]int{4})
}

func TestSum(t *testing.T) {
	assert.Equal(
		t,
		Sum([]int{1, 2, 3, 4}),
		10)
}

func TestTakeLast(t *testing.T) {
	assert.Equal(
		t,
		TakeLast([]int{1, 2, 3, 4}, 2),
		[]int{3, 4})
	assert.Panic(
		t,
		func() { TakeLast([]int{1, 2, 3, 4}, 5) },
		InvalidParam)
	assert.Panic(
		t,
		func() { TakeLast([]int{1, 2, 3, 4}, -1) },
		InvalidParam)
}

func TestTakeWhile(t *testing.T) {
	assert.Equal(
		t,
		TakeWhile([]int{1, 2, 3, 4}, func(x int) bool {
			return x%2 == 1
		}),
		[]int{1})
}

func TestToDictionary(t *testing.T) {
	dict := ToDictionary([]int{1, 2}, func(x int) string {
		return strconv.Itoa(x)
	})
	assert.Equal(t, len(dict), 2)
	assert.Equal(t, dict["1"], 1)
	assert.Equal(t, dict["2"], 2)
}

func TestLookup(t *testing.T) {
	lookup := ToLookUp(
		[]int{1, 2},
		func(x int) string {
			return strconv.Itoa(x)
		},
		func(x int) string {
			return strconv.Itoa(x * x)
		})
	assert.Equal(t, len(lookup), 2)
	assert.Equal(t, lookup["1"], "1")
	assert.Equal(t, lookup["2"], "4")
}

func TestUnion(t *testing.T) {
	assert.Equal(
		t,
		Order(Union([]int{1, 3, 4}, []int{1, 2, 5})),
		[]int{1, 2, 3, 4, 5})
}

func TestUnionBy(t *testing.T) {
	assert.Equal(
		t,
		Order(UnionBy([]int{1, 3, 4}, []int{1, 2, 5}, func(x int) string {
			return strconv.Itoa(x)
		})),
		[]int{1, 2, 3, 4, 5})
}

func TestZip(t *testing.T) {
	zip := Zip([]int{1, 2}, []string{"1", "2"})
	assert.Equal(t, len(zip), 2)
	assert.Equal(t, zip[0].First, 1)
	assert.Equal(t, zip[0].Second, "1")
	assert.Equal(t, zip[1].First, 2)
	assert.Equal(t, zip[1].Second, "2")
	assert.Panic(
		t,
		func() { Zip([]int{1}, []int{2, 3}) },
		InvalidParam)
}
